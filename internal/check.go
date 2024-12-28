package internal

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/sergi/go-diff/diffmatchpatch"
)

const (
	TestInputFileSuffix      = "_in"
	TestOutputFileSuffix     = "_out"
	TestExpectedFileSuffix   = "_exp"
	TestDebugFileSuffix      = "_debug"
	TestCasesDirectorySuffix = "_cases"

	MainSolutionName  = "main"
	LightSolutionName = "light"
)

func (pj *Project) Check(solution string, testsuite string) error {
	pj.logger.Debugf("starting check for solution \"%s\" with testsuit \"t%s\"", solution, testsuite)
	var execConfig *ExecConfig

	testsuiteDirName := testsuite + TestCasesDirectorySuffix

	switch solution {
	case MainSolutionName:
		execConfig = &pj.cfg.Main
	case LightSolutionName:
		if pj.cfg.Light != nil {
			execConfig = pj.cfg.Light
		} else {
			return errors.New("\"light\" is not initialized for this project ")
		}
	default:
		return fmt.Errorf("unknown solution: %s, only \"main\" or \"light\" allowed", solution)
	}

	if execConfig == nil {
		return fmt.Errorf("can't initialize checked command by solution name \"%s\"", solution)
	}

	pj.logger.Infof("Starting build command")
	err := runCmd(&execConfig.BuildCmd, nil, os.Stdout, os.Stderr, pj.logger)
	if err != nil {
		pj.logger.Errorf("Build Command failed: %s", err)
		return nil
	}

	pj.logger.Infof("Build Command done successfully")

	log.Debugf("Start porcessing testsuit %s", testsuite)
	testsuiteDirEntries, err := os.ReadDir(testsuiteDirName)
	if err != nil {
		log.Errorf("Failed to find info about testsuit dir %s: %s", testsuiteDirName, err)
	}

	testnames := make([]string, 0, len(testsuiteDirEntries)/2)

	for _, entry := range testsuiteDirEntries {
		log.Debugf("Found file %s", entry.Name())
		if !entry.Type().IsRegular() {
			log.Debugf("File %s is not regular skip it", entry.Name())
			continue
		}
		if !strings.HasSuffix(entry.Name(), TestInputFileSuffix) {
			log.Debugf("File %s name dosen't match test input pattern", entry.Name())
			continue
		}
		testname := strings.TrimSuffix(entry.Name(), TestInputFileSuffix)
		testnames = append(testnames, testname)
	}

	log.Debugf("found testcases: %s", testnames)

	for _, testname := range testnames {
		pj.logger.Infof("Processing test %s", coloredTestName(testname)) // TODO: add some color
		inputfileName := path.Join(testsuiteDirName, testname+TestInputFileSuffix)
		inpFile, err := os.Open(inputfileName)
		if err != nil {
			pj.logger.Errorf("Failed to open input file %s: %s", inputfileName, err)
			continue
		}
		defer inpFile.Close()

		inpFileContent, err := io.ReadAll(inpFile)
		if err != nil {
			pj.logger.Errorf("Read all input file failed: %s", err)
			return nil
		}

		pj.logger.Debugf("Starting Main Command")
		var outBld strings.Builder

		err = runCmd(&execConfig.RunCmd, strings.NewReader(string(inpFileContent)), &outBld, os.Stderr, pj.logger)
		if err != nil {
			pj.logger.Errorf("Main Command failed: %s", err)
			return nil
		}
		pj.logger.Debugf("Main command done successfully")

		expectedFileName := path.Join(testsuiteDirName, testname+TestOutputFileSuffix)
		expectedFile, err := os.Open(expectedFileName)

		var expectedOut string

		if err != nil {
			pj.logger.Errorf("Failed to open input file %s: %s", expectedFileName, err)
			pj.logger.Warn(coloredWarning())
			continue
		} else {
			defer expectedFile.Close()
			expectedOutB, err := io.ReadAll(expectedFile)
			if err != nil {
				pj.logger.Errorf("Failed to read output file: %s", err)
				pj.logger.Warn(coloredWarning())
				continue
			}
			expectedOut = string(expectedOutB)

		}

		outStr := outBld.String()
		// TODO: сделать отдельную нормализацию?
		outStr = strings.Trim(outStr, "\n ")

		expectedStr := expectedOut
		expectedStr = strings.Trim(expectedStr, "\n ")

		if expectedStr != outStr {
			pj.logger.Infof(coloredFailed())
			// TODO: add report on failed
			pj.logger.Infof("Input: \n%s\n", string(inpFileContent))
			pj.logger.Infof("Expected: \n%s\n", expectedStr)
			pj.logger.Infof("Got:\n%s\n\n", outStr)
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(expectedStr, outStr, false)
			pj.logger.Infof("Diff:\n %s \n \n", dmp.DiffPrettyText(diffs))
		} else {
			pj.logger.Info(coloredOk())
		}
	}

	return nil
}
