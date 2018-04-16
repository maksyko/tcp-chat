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
	"github.com/ievgen-ma/tcp-chat/protocol"
	"context"
	"log"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		ID := protocol.ID{
			Id:args[0],
		}
		m.RLock()
		id, err := client.Delete(context.Background(), &ID)
		if err != nil {
			return err
		}
		m.RUnlock()

		log.Println("result:", id)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
