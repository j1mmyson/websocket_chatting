package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type client struct{}

func Server() {
	app := fiber.New()
	stdin := bufio.NewReader(os.Stdin)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home.html", fiber.Map{})
	})

	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("middleware")
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		} else {
			fmt.Println("not websocket")
			return nil
		}
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		defer func() {
			c.Close()
		}()
		go func() {
			for {
				mt, msg, err := c.ReadMessage()
				if err != nil {
					fmt.Println("error read: ", err)
					break
				}
				log.Printf("recv: %s (type: %v)", msg, mt)
			}
		}()

		for {
			sendMessage := ""
			_, err := fmt.Scanf("%q", &sendMessage)
			if err != nil {
				fmt.Println("error while scanning message: ", err)
				stdin.ReadString('\n')
			}
			fmt.Printf("send: %s", sendMessage)
			if err = c.WriteMessage(1, []byte(sendMessage)); err != nil {
				fmt.Println("err while sending message: ", err)
				break
			}

		}

	}))

	app.Listen(":8080")
}
