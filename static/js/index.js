let popup = document.getElementById('popup');
let artists = [];

function openPopup(id) {
  const artist = artists.find(artist => artist.id === id);

  if (artist) {
      const img = popup.querySelector('img');
      const h2 = popup.querySelector('h2');
      const p = popup.querySelector('p');

      for (let [location, dates] of artist.relations) {
        p.textContent += `${location}\n`;
        h2.textContent += `${dates.join('\n')}`;
      }
      

      // cards.style.display = 'grid';
      // cards.style.gridTemplateColumns = 'repeat(4, 1fr)';
      // cards.style.gap = '2rem';
      // cards.style.maxWidth = 'var(--container-width)';
      // cards.style.margin = '0 auto';
      // cards.style.padding = 'var(--spacing-4)';
      // book.innerHTML = artist.relations.dates.map(date => `${date}`).join('<br>');
      // cover.innerHTML = artist.relations.location;
      img.src = artist.image;
      
      img.alt = artist.name;
      h2.textContent = `Thank you, ${artist.name}`;

      popup.classList.add('open-popup');
  } else {
      console.error('Artist not found for ID:', id);
  }
}

function closePopup() {
    popup.classList.remove('open-popup');
}
  
async function fetchArtists() {
  try {
      const response = await fetch('/api/artists');
      if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
      }
      artists = await response.json();
      for (let artist of artists) {
        artist.relations = new Map(Object.entries(artist.Relations));
      }
      console.log(artists);
  } catch (error) {
      console.error('Error fetching artists:', error);
  }
}

function initialize() {
  fetchArtists().then(fetchedArtists => {
    if (fetchedArtists && fetchedArtists.length > 0) {
      printArtistRelations(fetchedArtists[0].relations);
    }
  });
}

initialize();
  