let popup = document.getElementById('popup');
let artists = [];

function openPopup(id) {
  const artist = artists.find(artist => artist.id === id);

  if (artist) {
      const img = popup.querySelector('img');

      const locationGrid = document.querySelector('.location-grid');
      locationGrid.innerHTML = '';

      for (let [location, dates] of artist.relations) {
        const locationDiv = document.createElement('div');
        locationDiv.classList.add('location-item');
        
        const locationTitle = document.createElement('h3');
        locationTitle.textContent = location;
        
        const datesParagraph = document.createElement('p');
        datesParagraph.textContent = dates.join(', ');

        locationDiv.appendChild(locationTitle);
        locationDiv.appendChild(datesParagraph);
        locationGrid.appendChild(locationDiv);
      };
      
      img.src = artist.image;
      
      img.alt = artist.name;

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
  