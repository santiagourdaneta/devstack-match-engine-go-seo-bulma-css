// assets/js/main.js
document.addEventListener("DOMContentLoaded", () => {
    // Solo cargamos los observadores si el navegador lo soporta
    if ("IntersectionObserver" in window) {
        const imageObserver = new IntersectionObserver((entries, observer) => {
            entries.forEach(entry => {
                if (entry.isIntersecting) {
                    const image = entry.target;
                    image.src = image.dataset.src;
                    image.classList.remove("lazy");
                    imageObserver.unobserve(image);
                }
            });
        });

        document.querySelectorAll("img.lazy").forEach(img => imageObserver.observe(img));
    }
});