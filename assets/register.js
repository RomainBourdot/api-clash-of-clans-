const accountIcon = document.querySelector(".account-icon");
const accountLink = document.querySelector(".account-link");
const accountMenu = document.querySelector(".account-menu");

accountLink.addEventListener("mouseover", () => {
  accountIcon.style.stroke = "#80c342";
});

accountLink.addEventListener("mouseout", () => {
  accountIcon.style.stroke = "#ebec62";
});

accountLink.addEventListener("click", () => {
  accountMenu.classList.toggle("visible");
});

document.addEventListener("click", (event) => {
  if (
    !accountLink.contains(event.target) &&
    !accountMenu.contains(event.target)
  ) {
    accountMenu.classList.remove("visible");
  }
});