/*
   Copyright 2020 The Compose Specification Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (

    "io/ioutil"
    "C"

    "github.com/compose-spec/compose-go/loader"
    compose "github.com/compose-spec/compose-go/types"
)

//export load
func load(file string) *C.char {
    b, err := ioutil.ReadFile(file)
    if err != nil {
        return nil
    }
    config, err := loader.ParseYAML(b)
    if err != nil {
        return nil
    }
    var files []compose.ConfigFile
    files = append(files, compose.ConfigFile{Filename: file, Config: config})
    c, err := loader.Load(compose.ConfigDetails{
        WorkingDir:  ".",
        ConfigFiles: files,
    })
    if err != nil {
        return nil
    }
    enc, err := c.MarshalJSON()
     if err != nil {
        return nil
    }
    var l string
    l = string(enc)
    return C.CString(l)
}

func main() {}