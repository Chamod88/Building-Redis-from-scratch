# Discoveries & Learnings

Here is a running log of the key concepts and terms discovered while building Redis from scratch in Go.

### 1. In-Memory Store
* **What it means:** Storing information in the computer's "short-term memory" (RAM) instead of the hard drive.
* **Simple term:** Imagine keeping an important phone number in your head because it's super fast to recall, rather than walking over to a filing cabinet to look it up.

### 2. AOF (Append-Only File)
* **What it means:** A way to save data so it isn't lost if the computer turns off.
* **Simple term:** It’s like keeping a journal. Every time a change happens, you just write it at the very bottom of the page. If you ever forget what happened, you just read the journal from top to bottom to catch up.

### 3. RESP (Redis Serialization Protocol)
* **What it means:** The specific format of messages sent between the database and the user.
* **Simple term:** It's a strict set of grammar rules. It ensures that when you send a message, the database knows exactly how to read it without any confusion.

### 4. Mutex (or Lock)
* **What it means:** A tool to prevent errors when multiple parts of the program try to change data at the exact same time.
* **Simple term:** Think of it like a "talking stick" or a bathroom key. It ensures that only one person can change a piece of information at a time, preventing two people from writing over each other's work and making a mess.

### 5. Concurrency
* **What it means:** The ability of the program to handle multiple tasks at the same time.
* **Simple term:** It's like having multiple cashiers at a grocery store instead of just one. It allows our server to talk to dozens of users at the exact same time without making anyone wait in a long line.

### 6. TCP Socket
* **What it means:** The continuous connection line between a client (user) and the server.
* **Simple term:** It's like a dedicated telephone line between you and a friend. As long as neither of you hangs up, you can keep sending messages back and forth instantly.