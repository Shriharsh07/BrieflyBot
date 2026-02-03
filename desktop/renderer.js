const messages = document.getElementById("messages");

document.getElementById("minimize").onclick = () => {
  window.controls.minimize();
};

document.getElementById("close").onclick = () => {
  window.controls.close();
};

function addCard(title, text) {
  const card = document.createElement("div");
  card.className = "card";

  card.innerHTML = `
    <div class="card-title">${title}</div>
    <div class="card-body">${text}</div>
  `;

  messages.prepend(card);
}

// Mock data
addCard("🧠 Email Summary", "3 urgent emails · 1 invoice due today");
addCard("📅 Reminder", "Rent follow-up scheduled for tomorrow");
addCard("📌 Task", "Follow up with tenant regarding payment");
addCard("🧠 Email Summary", "3 urgent emails · 1 invoice due today");
addCard("📅 Reminder", "Rent follow-up scheduled for tomorrow");
addCard("📌 Task", "Follow up with tenant regarding payment");
addCard("🧠 Email Summary", "3 urgent emails · 1 invoice due today");
addCard("📅 Reminder", "Rent follow-up scheduled for tomorrow");
addCard("📌 Task", "Follow up with tenant regarding payment");

