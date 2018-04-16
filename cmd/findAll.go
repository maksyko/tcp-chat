// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"

	"context"
	"github.com/ievgen-ma/tcp-chat/protocol"
	"log"
)

// findAllCmd represents the findAll command
var findAllCmd = &cobra.Command{
	Use:   "find_all",
	Short: "Find all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		m.RLock()
		void := protocol.Void{}
		tasks, err := client.FindAll(context.Background(), &void)
		if err != nil {
			return err
		}
		m.RUnlock()

		log.Println("result:", tasks)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(findAllCmd)

}
