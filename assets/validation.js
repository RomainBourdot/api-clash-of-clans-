document.addEventListener('DOMContentLoaded', function() {
  const searchForm = document.getElementById('search-form');

  searchForm.addEventListener('submit', function(e) {
    const searchInput = document.getElementById('search-bar');
    const errorDiv = document.getElementById('search-error');
    const value = searchInput.value.trim();

    // Vérifier si le champ est vide
    if (value === "") {
      e.preventDefault();
      errorDiv.textContent = "La recherche ne peut pas être vide";
      return;
    }

    // Vérifier si le contenu est uniquement composé de chiffres
    if (/^\d+$/.test(value)) {
      e.preventDefault();
      errorDiv.textContent = "La recherche ne peut pas être un chiffre uniquement";
      return;
    }
    
    // Compter le nombre de lettres (a-z, A-Z et lettres accentuées)
    const letterCount = value.replace(/[^a-zA-ZÀ-ÿ]/g, "").length;
    if (letterCount < 3) {
      e.preventDefault();
      errorDiv.textContent = "La recherche doit contenir au moins trois lettres";
      return;
    }
    
    // Si tout est OK, on efface le message d'erreur
    errorDiv.textContent = "";
  });
});
