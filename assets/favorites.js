document.addEventListener("DOMContentLoaded", function() {
    var el = document.getElementById('favorites-list');
    if (el) {
      var sortable = Sortable.create(el, {
        animation: 150
      });
  
      var saveOrderBtn = document.getElementById('saveOrder');
      if (saveOrderBtn) {
        saveOrderBtn.addEventListener('click', function() {
          var orderedTags = [];
          el.querySelectorAll('li').forEach(function(li) {
            orderedTags.push(li.getAttribute('data-tag'));
          });
  
          fetch('/favorites/reorder', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify(orderedTags)
          })
          .then(response => {
            if (response.ok) {
              alert('Ordre mis à jour avec succès');
            } else {
              alert('Erreur lors de la mise à jour de l\'ordre');
            }
          });
        });
      }
    }
  });
  