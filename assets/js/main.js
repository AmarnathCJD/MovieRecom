const menuBtn = document.getElementById('mobile-menu-button');
const navLinks = document.getElementById('menu');

menuBtn.addEventListener('click', () => {
  navLinks.classList.toggle('hidden');
});

const searchBtn = document.getElementById('search-button');
const searchInput = document.getElementById('search-input');
const searchResults = document.getElementById('search-results');
const searchGrid = document.getElementById('grid');

function initSearchGrid() {
  const width = window.innerWidth;
  if (width < 768) {
    if (searchGrid.classList.contains('grid-cols-6')) {
      searchGrid.classList.remove('grid-cols-6');
    }
    if (!searchGrid.classList.contains('grid-cols-2')) {
      searchGrid.classList.add('grid-cols-2');
    }
  } else {
    if (searchGrid.classList.contains('grid-cols-2')) {
      searchGrid.classList.remove('grid-cols-2');
    }
    if (!searchGrid.classList.contains('grid-cols-6')) {
      searchGrid.classList.add('grid-cols-6');
    }
  }
}

window.onresize = () => {
  initSearchGrid();
}

initSearchGrid();

searchInput.addEventListener('keyup', (e) => {
  const val = e.target.value;
  if (val.length % 2 == 0 && val.length > 4) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', '/api/search?q=' + val, true);
    xhr.onload = function () {
      if (this.status == 200) {
        const data = JSON.parse(this.responseText);
        searchGrid.innerHTML = '';
        data.forEach((item) => {
          var poster = item.poster;
          if (poster != undefined) {
            if (item.imdb_rating == undefined) {
              item.imdb_rating = 0;
            }
            const div = document.createElement('div');
            poster = item.poster.replace('840x', '210x');
            div.classList.add('bg-white', 'rounded-lg', 'shadow-md', 'mx-2', 'w-32');
            div.innerHTML = `
          <img src="${poster}" alt='${item.name}' loading="lazy"
              class="h-48 object-cover rounded-t-lg">
          <div class="p-2">
              <h2 class="text-sm font-semibold text-gray-800">${correctName(item.name)}</h2>
              <div class="flex items-center mt-1">
                  <svg class="h-4 w-4 fill-current text-yellow-500" viewBox="0 0 24 24">
                      <path
                          d="M12 17.27l5.74 3.28l-1.1-6.43l4.64-4.52l-6.42-.94L12 3.05L9.14 8.06l-6.42.94l4 4.52l-1.1 6.43L12 17.27z">
                      </path>
                  </svg>
                  <span class="text-xs ml-1 text-gray-700">${item.imdb_rating}</span>
              </div>
          </div>
          `;
            searchGrid.appendChild(div);

            if (searchResults.classList.contains('hidden')) {
              searchResults.classList.remove('hidden');
            }

            div.addEventListener('mouseover', () => {
              div.classList.add('border-2');
              div.classList.add('border-yellow-500');
              div.classList.add('shadow-lg');
            });

            div.addEventListener('mouseout', () => {
              div.classList.remove('border-2');
              div.classList.remove('border-yellow-500');
              div.classList.remove('shadow-lg');
            });
          }
        });
      }
    }

    xhr.send();
  }
})

function correctName(name) {
  const maxMin = 40;
  if (name.length > maxMin) {
    return name.substring(0, maxMin);
  } else if (name.length < maxMin) {
    name += ' '.repeat(maxMin - name.length);
    return name + '<br>';
  } else {
    return name;
  }
}
