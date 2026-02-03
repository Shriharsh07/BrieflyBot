const messages = document.getElementById("messages");

/* Window buttons */
document.getElementById("minimize").addEventListener("click", () => {
  window.windowControls.minimize();
});

document.getElementById("close").addEventListener("click", () => {
  window.windowControls.close();
});

/* Cards */
function addCard(title, body) {
  const card = document.createElement("div");
  card.className = "card";
  card.innerHTML = `
    <div class="card-title">${title}</div>
    <div class="card-body">${body}</div>
  `;
  messages.prepend(card);
}

/* WebSocket */
window.wsClient.connect(
  "ws://localhost:8080/ws",
  (data) => {
    addCard(data.title || "Update", data.body || "");
  },
  (status) => {
    addCard(
      status === "connected" ? "🟢 Connected" : "🔴 Disconnected",
      `WebSocket status: ${status}`
    );
  }
);

/* Mock cards */
addCard("🧠 Email Summary", "3 urgent emails · 1 invoice due today");
addCard("📅 Reminder", "Rent follow-up scheduled for tomorrow");
addCard("📌 Task", "Follow up with tenant regarding payment");
