package copilot

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"time"
)

type Pong struct {
	Now    int    `json:"now"`
	Status string `json:"status"`
	Ns1    string `json:"ns1"`
}

func _ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Pong{
		Now:    time.Now().Second(),
		Status: "ok",
		Ns1:    "200 OK",
	})
}

func models(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": []gin.H{
			{
				"capabilities": gin.H{
					"family":    "gpt-3.5-turbo",
					"limits":    gin.H{"max_prompt_tokens": 12288},
					"object":    "model_capabilities",
					"supports":  gin.H{"tool_calls": true},
					"tokenizer": "cl100k_base",
					"type":      "chat",
				},
				"id":      "gpt-3.5-turbo",
				"name":    "GPT 3.5 Turbo",
				"object":  "model",
				"version": "gpt-3.5-turbo-0613",
			},
			{
				"capabilities": gin.H{
					"family":    "gpt-3.5-turbo",
					"limits":    gin.H{"max_prompt_tokens": 12288},
					"object":    "model_capabilities",
					"supports":  gin.H{"tool_calls": true},
					"tokenizer": "cl100k_base",
					"type":      "chat",
				},
				"id":      "gpt-3.5-turbo-0613",
				"name":    "GPT 3.5 Turbo",
				"object":  "model",
				"version": "gpt-3.5-turbo-0613",
			},
			{
				"capabilities": gin.H{
					"family":    "gpt-4o-mini",
					"limits":    gin.H{"max_prompt_tokens": 12288},
					"object":    "model_capabilities",
					"supports":  gin.H{"tool_calls": true, "parallel_tool_calls": true},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":      "gpt-4o-mini",
				"name":    "GPT 4o Mini",
				"object":  "model",
				"version": "gpt-4o-mini-2024-07-18",
			},
			{
				"capabilities": gin.H{
					"family":    "gpt-4o-mini",
					"limits":    gin.H{"max_prompt_tokens": 12288},
					"object":    "model_capabilities",
					"supports":  gin.H{"tool_calls": true, "parallel_tool_calls": true},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":      "gpt-4o-mini",
				"name":    "GPT 4o Mini",
				"object":  "model",
				"version": "gpt-4o-mini-2024-07-18",
			},
			{
				"capabilities": gin.H{
					"family":    "gpt-4",
					"limits":    gin.H{"max_prompt_tokens": 20000},
					"object":    "model_capabilities",
					"supports":  gin.H{"tool_calls": true},
					"tokenizer": "cl100k_base",
					"type":      "chat",
				},
				"id":      "gpt-4",
				"name":    "GPT 4",
				"object":  "model",
				"version": "gpt-4-0613",
			},
			{
				"capabilities": gin.H{
					"family":    "gpt-4",
					"limits":    gin.H{"max_prompt_tokens": 20000},
					"object":    "model_capabilities",
					"supports":  gin.H{"tool_calls": true},
					"tokenizer": "cl100k_base",
					"type":      "chat",
				},
				"id":      "gpt-4-0613",
				"name":    "GPT 4",
				"object":  "model",
				"version": "gpt-4-0613",
			},
			{
				"capabilities": gin.H{
					"family":    "gpt-4-turbo",
					"limits":    gin.H{"max_prompt_tokens": 20000},
					"object":    "model_capabilities",
					"supports":  gin.H{"parallel_tool_calls": true, "tool_calls": true},
					"tokenizer": "cl100k_base",
					"type":      "chat",
				},
				"id":      "gpt-4-0125-preview",
				"name":    "GPT 4 Turbo",
				"object":  "model",
				"version": "gpt-4-0125-preview",
			},
			{
				"capabilities": gin.H{
					"family":    "gpt-4o",
					"limits":    gin.H{"max_prompt_tokens": 20000},
					"object":    "model_capabilities",
					"supports":  gin.H{"parallel_tool_calls": true, "tool_calls": true},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":      "gpt-4o",
				"name":    "GPT 4o",
				"object":  "model",
				"version": "gpt-4o-2024-05-13",
			},
			{
				"capabilities": gin.H{
					"family":    "gpt-4o",
					"limits":    gin.H{"max_prompt_tokens": 20000},
					"object":    "model_capabilities",
					"supports":  gin.H{"parallel_tool_calls": true, "tool_calls": true},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":      "gpt-4o-2024-05-13",
				"name":    "GPT 4o",
				"object":  "model",
				"version": "gpt-4o-2024-05-13",
			},
			{
				"capabilities": gin.H{
					"family":    "gpt-4o",
					"limits":    gin.H{"max_prompt_tokens": 20000},
					"object":    "model_capabilities",
					"supports":  gin.H{"parallel_tool_calls": true, "tool_calls": true},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":      "gpt-4-o-preview",
				"name":    "GPT 4o",
				"object":  "model",
				"version": "gpt-4-o-preview",
			},
			{
				"capabilities": gin.H{
					"family":    "text-embedding-ada-002",
					"limits":    gin.H{"max_inputs": 256},
					"object":    "model_capabilities",
					"supports":  gin.H{},
					"tokenizer": "cl100k_base",
					"type":      "embeddings",
				},
				"id":      "text-embedding-ada-002",
				"name":    "Embedding V2 Ada",
				"object":  "model",
				"version": "text-embedding-ada-002",
			},
			{
				"capabilities": gin.H{
					"family":    "text-embedding-3-small",
					"limits":    gin.H{"max_inputs": 256},
					"object":    "model_capabilities",
					"supports":  gin.H{"dimensions": true},
					"tokenizer": "cl100k_base",
					"type":      "embeddings",
				},
				"id":      "text-embedding-3-small",
				"name":    "Embedding V3 small",
				"object":  "model",
				"version": "text-embedding-3-small",
			},
			{
				"capabilities": gin.H{
					"family":    "text-embedding-3-small",
					"object":    "model_capabilities",
					"supports":  gin.H{"dimensions": true},
					"tokenizer": "cl100k_base",
					"type":      "embeddings",
				},
				"id":      "text-embedding-3-small-inference",
				"name":    "Embedding V3 small (Inference)",
				"object":  "model",
				"version": "text-embedding-3-small",
			},
		},
		"object": "list",
	})
}

func closeIO(c io.Closer) {
	err := c.Close()
	if nil != err {
		log.Println(err)
	}
}
