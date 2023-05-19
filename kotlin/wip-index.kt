import io.ktor.application.*
import io.ktor.features.ContentNegotiation
import io.ktor.features.StatusPages
import io.ktor.http.HttpStatusCode
import io.ktor.jackson.jackson
import io.ktor.request.receive
import io.ktor.response.respond
import io.ktor.routing.*
import io.ktor.server.engine.embeddedServer
import io.ktor.server.netty.Netty

data class Note(val id: String, val title: String, val description: String)

val notes = mutableListOf<Note>()

fun generateId(): String {
    val idLength = 9
    val characters = ('a'..'z') + ('A'..'Z') + ('0'..'9')
    return (1..idLength)
        .map { characters.random() }
        .joinToString("")
}

fun Application.module() {
    install(ContentNegotiation) {
        jackson { }
    }

    install(StatusPages) {
        exception<Throwable> { cause ->
            call.respond(HttpStatusCode.InternalServerError, cause.localizedMessage)
        }
    }

    routing {
        route("/notes") {
            post {
                val note = call.receive<Note>()
                val id = generateId()
                val newNote = note.copy(id = id)
                notes.add(newNote)
                call.respond(HttpStatusCode.Created, newNote)
            }

            get {
                call.respond(notes)
            }

            route("/{id}") {
                get {
                    val id = call.parameters["id"]
                    val note = notes.find { it.id == id }
                    if (note != null) {
                        call.respond(note)
                    } else {
                        call.respond(HttpStatusCode.NotFound, "Note not found")
                    }
                }

                put {
                    val id = call.parameters["id"]
                    val note = call.receive<Note>()
                    val existingNote = notes.find { it.id == id }
                    if (existingNote != null) {
                        val updatedNote = existingNote.copy(
                            title = note.title,
                            description = note.description
                        )
                        notes.remove(existingNote)
                        notes.add(updatedNote)
                        call.respond(updatedNote)
                    } else {
                        call.respond(HttpStatusCode.NotFound, "Note not found")
                    }
                }

                delete {
                    val id = call.parameters["id"]
                    val note = notes.find { it.id == id }
                    if (note != null) {
                        notes.remove(note)
                        call.respond(note)
                    } else {
                        call.respond(HttpStatusCode.NotFound, "Note not found")
                    }
                }
            }
        }
    }
}

fun main() {
    embeddedServer(Netty, port = 8080, module = Application::module).start(wait = true)
}
