package handler

import (
	"os"
	"time"
	"sort"
	"strconv"
	"net/http"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"tutorials/file"
	"github.com/gin-gonic/gin"
)

func ListContent(c *gin.Context) {
	tutorial := c.Param("tutorial")
	c.JSON(http.StatusOK, getContents(tutorial))	
}

func InsertContent(c *gin.Context) {
	number, err := strconv.Atoi(c.PostForm("number"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Número inválido"})
		return
	}
	
	tutorial := c.PostForm("tutorial")
	title := c.PostForm("title")
	scontent := c.PostForm("content")
	code := c.PostForm("code")
	jsonPath := ROOT_DIR + "data/" + tutorial + ".json"

	now := time.Now()
	hash := md5.Sum([]byte(title + scontent + code + now.Format("yyyymmddhhmmss")))
	id := hex.EncodeToString(hash[:])

	commands := getContents(tutorial)

	command := file.Command{
		Number: number,
		Title: title,
	}

	content := file.Content{
		ID: id,
		Content: scontent,
		Code: code,
	}

	exist := false
	for i := range commands {
		if commands[i].Title == command.Title {
			exist = true
			commands[i].Content = append(commands[i].Content, content)
			break
		}
	}

	if !exist {
		command.Content = append(command.Content, content)
		commands = append(commands, command)
	}

	sort.Slice(commands, func(i, j int) bool {
    return commands[i].Number < commands[j].Number
	})

	file.SaveCommand(jsonPath, commands)

	c.JSON(http.StatusOK, gin.H{"message": "Conteúdo salvo com sucesso."})
}

func UpdateContent(c *gin.Context) {
	number, err := strconv.Atoi(c.PostForm("number"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Número inválido"})
		return
	}
	
	tutorial := c.PostForm("tutorial")
	oldTitle := c.PostForm("oldTitle")
	title := c.PostForm("title")
	id := c.PostForm("id")
	content := c.PostForm("content")
	code := c.PostForm("code")
	jsonPath := ROOT_DIR + "data/" + tutorial + ".json"

	commands := getContents(tutorial)

	for i := range commands {
		if commands[i].Title == oldTitle {
			commands[i].Number = number
			commands[i].Title = title
			for j := range commands[i].Content {
				if commands[i].Content[j].ID == id {
					commands[i].Content[j].Content = content
					commands[i].Content[j].Code = code
					break
				}
			}
			break
		}
	}

	sort.Slice(commands, func(i, j int) bool {
    return commands[i].Number < commands[j].Number
	})

	file.SaveCommand(jsonPath, commands)

	c.JSON(http.StatusOK, gin.H{"message": "Conteúdo alterado com sucesso."})
}

func DeleteContent(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	tutorial := c.Param("tutorial")
	filePath := ROOT_DIR + "data/" + tutorial + ".json"

	commands := getContents(tutorial)

	for i := range commands {
		if commands[i].Title == title {
			for j := range commands[i].Content {
				if commands[i].Content[j].ID == id {
					commands[i].Content = append(commands[i].Content[:j], commands[i].Content[j+1:]...)
					break
				}
			}

			if len(commands[i].Content) < 1 {
				commands = append(commands[:i], commands[i+1:]...)
				break
			}
		}
	}

	file.SaveCommand(filePath, commands)

	c.JSON(http.StatusOK, gin.H{"message": "Conteúdo apagado com sucesso."})
}

func getContents(tutorial string) []file.Command {
	filePath := ROOT_DIR + "data/" + tutorial + ".json"
	var jsonData []file.Command
	
	file, err := os.ReadFile(filePath)
	if err != nil {
		println("Erro: ", err.Error())
		return jsonData
	}

	err = json.Unmarshal(file, &jsonData)
	if err != nil {
		println("Erro: ", err.Error())
		return jsonData
	}

	return jsonData
}