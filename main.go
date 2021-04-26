package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/jessevdk/go-flags"
)

const (
	cmdName = "getx-generate"
)

var (
	options Options
	parser  *flags.Parser
)

type Options struct {
	Output string `short:"o" long:"output" description:"output path" default:"lib/modules"`
}

func init() {
	parser = flags.NewParser(&options, flags.Default)
	parser.Name = cmdName
	parser.Usage = "[OPTIONS] name"
}

func main() {
	args, err := parser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); ok && e.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			exitOnErrorWriteHelp("")
		}
	}

	if len(args) == 1 {
		name := args[0]
		outPath := options.Output
		if !dirExists(outPath) {
			exitOnError(fmt.Sprintf("%v not exist", outPath))
		}
		dirPath := path.Join(outPath, name)
		if dirExists(dirPath) {
			exitOnError(fmt.Sprintf("%v is exist", name))
		}
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			exitOnError(err.Error())
		}
		bindingPath := path.Join(dirPath, fmt.Sprintf("%s_binding.dart", name))
		controllerPath := path.Join(dirPath, fmt.Sprintf("%s_controller.dart", name))
		pagePath := path.Join(dirPath, fmt.Sprintf("%s_page.dart", name))

		nameCamel := camelName(name)

		bindingStr := []byte(fmt.Sprintf(
			`import 'package:get/get.dart';
import '%s_controller.dart';

class %sBinding implements Bindings {
  @override
  void dependencies() {
    Get.lazyPut<%sController>(() => %sController());
  }
}
		`, name, nameCamel, nameCamel, nameCamel))
		ioutil.WriteFile(bindingPath, bindingStr, 0644)

		controllerStr := []byte(fmt.Sprintf(
			`import 'package:get/get.dart';

class %sController extends GetxController {}
`, nameCamel))
		ioutil.WriteFile(controllerPath, controllerStr, 0644)

		pageStr := []byte(fmt.Sprintf(
			`import 'package:flutter/material.dart';
import '%s_controller.dart';
import 'package:get/get.dart';

class %sPage extends GetView<%sController> {
  @override
  Widget build(Object context) {
    return Container();
  }
}
		`, name, nameCamel, nameCamel))
		ioutil.WriteFile(pagePath, pageStr, 0644)
	} else {
		exitOnErrorWriteHelp("Nothing to do...")
	}
}

func dirExists(name string) bool {
	if fi, err := os.Stat(name); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}
	return false
}

func camelName(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

func exitOnError(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func exitOnErrorWriteHelp(msg string) {
	fmt.Println(msg)
	parser.WriteHelp(os.Stderr)
	os.Exit(1)
}
