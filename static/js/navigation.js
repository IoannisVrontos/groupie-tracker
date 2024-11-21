document.querySelectorAll(".card__content").forEach((card) => {
  card.addEventListener("click", () => {
      // Get the ID from the `data-id` attribute
      const artistID = card.getAttribute("data-id");

      if (!artistID) {
          console.error("No ID found for navigation.");
          return;
      }

      // Construct the target URL and navigate
      const targetPageURL = `http://localhost:8080/artist/${artistID}`;
      window.location.href = targetPageURL; // Redirect to the new page
  });
});