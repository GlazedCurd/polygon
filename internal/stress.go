package internal

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/charmbracelet/log"
)

const (
	generatedCasesDirName    = "generated"
	stressTestWorkers        = 10
	casesDirFilesPermissions = 0777
	seedPlaceholder          = "{seed}"
)

type caseDescription struct {
	Seed uint64
	Num  uint64
}

func (pj *Project) Stress(seed uint64, number uint64) error {
	err := os.Mkdir(generatedCasesDirName+TestCasesDirectorySuffix, casesDirFilesPermissions)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("create test directory %w", err)
	}

	rand := rand.New(rand.NewSource(int64(seed)))

	seeds := make(chan caseDescription)

	// TODO: сделать проверку что верно проинициализированы
	// - команда генерации инпута
	// - команда запуска для лёгкого решения
	// - команда запуска для главного решения

	pj.logger.Infof("Building light solution")

	err = runCmd(&pj.cfg.Light.BuildCmd, nil, os.Stdout, os.Stderr, pj.logger)
	if err != nil {
		pj.logger.Errorf("Building Light Command failed: %s", err)
		return nil
	}
	pj.logger.Info("Light solution build stage finished")

	pj.logger.Infof("Building main solution")

	err = runCmd(&pj.cfg.Main.BuildCmd, nil, os.Stdout, os.Stderr, pj.logger)
	if err != nil {
		pj.logger.Errorf("Building Light Command failed: %s", err)
		return nil
	}

	pj.logger.Info("Main solution build stage finished")

	pj.logger.Info("Starting stress test")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := uint64(0); i < number; i++ {
			seeds <- caseDescription{
				Seed: rand.Uint64(),
				Num:  i,
			}
			if i%100 == 0 {
				pj.logger.Infof("Processed %d", i)
			}
		}
		close(seeds)
	}()

	for i := 0; i < stressTestWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for seed := range seeds {
				err := pj.runStressTestWithSeed(&seed)
				if err != nil {
					log.Errorf("Error on launching %s", err)
				}
			}
		}()
	}

	wg.Wait()

	pj.logger.Info("Stress test completed")

	return nil
}

func (pj *Project) dumpTestCases(testSeed *caseDescription, input string, output string, expexted string, stderr string) error {
	testname := fmt.Sprintf("%d_%d", testSeed.Seed%100, testSeed.Num)
	testdirname := generatedCasesDirName + TestCasesDirectorySuffix

	dumpFile := func(filename string, data string) error {
		fullFileName := path.Join(testdirname, filename)
		pj.logger.Debugf("dumping to file %s", fullFileName)
		dumpfile, err := os.OpenFile(fullFileName, os.O_WRONLY|os.O_CREATE, casesDirFilesPermissions)
		if err != nil {
			return fmt.Errorf("can't open file %w", err)
		}
		defer dumpfile.Close()

		_, err = dumpfile.WriteString(data)
		if err != nil {
			return fmt.Errorf("writing input to file %w", err)
		}
		return nil
	}

	var resError error

	err := dumpFile(testname+TestInputFileSuffix, input)
	if err != nil {
		log.Errorf("writing input file %s", err)
	}

	resError = errors.Join(resError, err)

	err = dumpFile(testname+TestOutputFileSuffix, output)
	if err != nil {
		log.Errorf("writing output file %s", err)
	}

	resError = errors.Join(resError, err)

	err = dumpFile(testname+TestExpectedFileSuffix, expexted)
	if err != nil {
		log.Errorf("writing expected file %s", err)
	}

	resError = errors.Join(resError, err)

	if strings.Trim(stderr, " \n\t") == "" {
		err = dumpFile(testname+TestDebugFileSuffix, stderr)
		if err != nil {
			log.Errorf("writing debug file %s", err)
		}
	}

	resError = errors.Join(resError, err)

	return resError
}

func (pj *Project) runStressTestWithSeed(testSeed *caseDescription) error {
	inp := CommandCfg{}
	var testLocalErrLog strings.Builder

	inp = pj.cfg.InputGenerator

	inp.Args = strings.ReplaceAll(inp.Args, seedPlaceholder, fmt.Sprint(testSeed.Seed))

	runner := func(cmd *CommandCfg, input io.Reader) (string, string, error) {
		var testLocalErrLog strings.Builder
		var cmdOutput strings.Builder
		err := runCmd(cmd, input, &cmdOutput, &testLocalErrLog, pj.logger)
		if err != nil {
			return "", testLocalErrLog.String(), fmt.Errorf("running test command %w", err)
		}
		return cmdOutput.String(), testLocalErrLog.String(), nil
	}

	testLocalErrLog.WriteString("input stderr\n")

	inputRunRes, errorLogRes, err := runner(&inp, nil)
	if err != nil {
		return fmt.Errorf("generating input %w", err)
	}

	if errorLogRes != "" {
		testLocalErrLog.WriteString("\n\n\n expect sol stderr:\n")
		testLocalErrLog.WriteString(errorLogRes)
	}

	expected, errorLogRes, err := runner(&pj.cfg.Light.RunCmd, strings.NewReader(inputRunRes))
	if err != nil {
		return fmt.Errorf("building expected result %w", err)
	}

	if errorLogRes != "" {
		testLocalErrLog.WriteString("\n\n\n expected sol stderr:\n")
		testLocalErrLog.WriteString(errorLogRes)
	}

	actual, errorLogRes, err := runner(&pj.cfg.Main.RunCmd, strings.NewReader(inputRunRes))
	if err != nil {
		return fmt.Errorf("building actual result %w", err)
	}

	if errorLogRes != "" {
		testLocalErrLog.WriteString("\n\n\n actual sol stderr:\n")
		testLocalErrLog.WriteString(errorLogRes)
	}

	if strings.TrimSpace(expected) != strings.TrimSpace(actual) {
		return pj.dumpTestCases(testSeed, inputRunRes, actual, expected, testLocalErrLog.String())
	}

	return nil
}
