const filterToggle = document.getElementById('filter-toggle');
const filters = document.getElementById('filters');

filterToggle.addEventListener('click', () => {
  if (filters.style.display === 'none' || filters.style.display === '') {
    filters.style.display = 'block';
  } else {
    filters.style.display = 'none';
  }
});