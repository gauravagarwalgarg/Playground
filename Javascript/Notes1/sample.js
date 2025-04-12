const btn = document.getElementById("toggleDarkMode");

btn.addEventListener("click", () => {
  document.body.classList.toggle("dark-mode");
});

