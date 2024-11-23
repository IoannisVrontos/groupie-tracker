let popup = document.getElementById('popup');
let artists = [];

function openPopup(id) {
  // Adjust the logic if necessary to handle the incremented ID
  const artist = artists.find(artist => artist.id === id); // Adjust if needed

  if (artist) {
      const img = popup.querySelector('img');
      const h2 = popup.querySelector('h2');
      const p = popup.querySelector('p');

      img.src = artist.image;
      img.alt = artist.name;
      h2.textContent = `Thank you, ${artist.name}`;
      p.textContent = `Info about ${artist.name}`;

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
  } catch (error) {
      console.error('Error fetching artists:', error);
  }
}

function initialize() {
  fetchArtists();
}

initialize();