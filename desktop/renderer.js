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

  let contentHTML = "";

  // If body contains line breaks → treat as bullet points
  if (body.includes("\n")) {
    const points = body
      .split("\n")
      .map(p => p.trim())
      .filter(Boolean);

    contentHTML = `
      <ul class="card-list">
        ${points.map(p => `<li>${p}</li>`).join("")}
      </ul>
    `;
  } else {
    contentHTML = `<div class="card-body">${body}</div>`;
  }

  card.innerHTML = `
    <div class="card-title">${title}</div>
    ${contentHTML}
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
addCard(
  "🧠 Email Summary",
  "3 urgent emails\n1 invoice due today\n2 newsletters skipped"
);

