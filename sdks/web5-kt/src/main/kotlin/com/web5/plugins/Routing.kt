package com.web5.plugins

import com.web5.plugins.routes.credentialIssue
import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import io.ktor.server.engine.stop
import kotlinx.coroutines.launch
import kotlin.system.exitProcess

fun Application.configureRouting() {
    routing {
        get("/") {
            val serverID = mapOf(
                "name" to "web5-kt",
                "language" to "Kotlin",
                "url" to "https://github.com/TBD54566975/web5-kt"
            )
            call.respond(serverID)
        }

        post("/credentials/issue") {
            call.credentialIssue()
        }

        get("/shutdown") {
            launch {
                exitProcess(0)
            }
            call.respondText("shutting down server..")
        }
    }
}
