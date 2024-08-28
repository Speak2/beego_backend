package middleware

import (
    "github.com/beego/beego/v2/server/web/context" // Corrected import path for Beego v2
)

// CORSMiddleware handles CORS settings
func CORSMiddleware(ctx *context.Context) {
    ctx.Output.Header("Access-Control-Allow-Origin", "*") // Allows all origins; you can restrict to specific origins if needed
    ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
    ctx.Output.Header("Access-Control-Allow-Credentials", "true")

    // If it's an OPTIONS request, terminate early with status 204
    if ctx.Input.Method() == "OPTIONS" {
        ctx.ResponseWriter.WriteHeader(204)
        return
    }
}
