document.addEventListener('DOMContentLoaded', function() {
  const searchForm = document.getElementById('search-form');

  searchForm.addEventListener('submit', function(e) {
    const searchInput = document.getElementById('search-bar');
    const errorDiv = document.getElementById('search-error');
    const value = searchInput.value.trim();

    if (value === "") {
      e.preventDefault();
      errorDiv.textContent = "La recherche ne peut pas être vide";
      return;
    }

    if (/^\d+$/.test(value)) {
      e.preventDefault();
      errorDiv.textContent = "La recherche ne peut pas être un chiffre uniquement";
      return;
    }
    
    const letterCount = value.replace(/[^a-zA-ZÀ-ÿ]/g, "").length;
    if (letterCount < 3) {
      e.preventDefault();
      errorDiv.textContent = "La recherche doit contenir au moins trois lettres";
      return;
    }
    
    errorDiv.textContent = "";
  });
});
