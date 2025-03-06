const filterToggle = document.getElementById('filter-toggle');
const filters = document.getElementById('filters');

// Au clic, alterne l'affichage du bloc
filterToggle.addEventListener('click', () => {
  if (filters.style.display === 'none' || filters.style.display === '') {
    filters.style.display = 'block';
  } else {
    filters.style.display = 'none';
  }
});