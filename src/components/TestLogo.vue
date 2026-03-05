<template>
  <div class="logo-container">
    <div class="spinning-border"></div>
    <div class="logo-wrapper">
      <img :src="logo" alt="Logo" class="logo-img" />
    </div>
  </div>
</template>

<script>
import logoImage from "@/assets/logopng.png";

export default {
  name: "TestLogo",
  data() {
    return {
      logo: logoImage,
    };
  },
};
</script>

<style scoped>
.logo-container {
  position: relative;
  width: 150px;
  height: 150px;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 50%;

  overflow: visible;
}

/* The colorful spinning background */
.spinning-border {
  position: absolute;
  top: -4px;
  left: -4px;
  right: -4px;
  bottom: -4px;
  border-radius: 50%;
  /* Default: White light pulsing */
  background: conic-gradient(
    from 0deg,
    transparent,
    rgba(255, 255, 255, 0.8),
    transparent
  );
  animation: pulse 2s ease-in-out infinite;
  /* Pulsing instead of spinning */
  z-index: 0;
  filter: blur(8px);
  /* Creates a glow effect */
  opacity: 0.9;
  transition: background 0.3s ease, opacity 0.3s ease;
}

/* A second sharp border layer if we want distinct lines */
.spinning-border::after {
  content: "";
  position: absolute;
  top: 4px;
  left: 4px;
  right: 4px;
  bottom: 4px;
  border-radius: 50%;
  background: conic-gradient(from 0deg, transparent, white, transparent);
  animation: pulse 2s ease-in-out infinite;
  z-index: 1;
  filter: blur(0px);
  transition: background 0.3s ease;
}

/* The wrapper creates the background for the logo to sit on, masking the center */
.logo-wrapper {
  position: relative;
  width: 100%;
  height: 100%;
  background-color: white;
  /* Match your page background */
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 2;
  /* Sit on top of the spinner */
  box-shadow: inset 0 0 20px rgba(0, 0, 0, 0.1);
}

.logo-img {
  width: 80%;
  /* Adjust size inside the circle */
  height: auto;
  cursor: pointer;
}

/* Hover Effects */
/* .logo-wrapper:hover .logo-img {  } */

.logo-container:hover .spinning-border,
.logo-container:hover .spinning-border::after {
  animation: spin 3s linear infinite;
  /* Spin continuously on hover, slower speed */
  opacity: 1;
  /* Switch to colorful gradient on hover */
  background: conic-gradient(
    from 0deg,
    #ff0000,
    #ff00ff,
    #00ffff,
    #00ff00,
    #ffff00,
    #ff0000
  );
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}

@keyframes pulse {
  0% {
    opacity: 0.5;
    transform: scale(0.98);
  }

  50% {
    opacity: 1;
    transform: scale(1.02);
  }

  100% {
    opacity: 0.5;
    transform: scale(0.98);
  }
}
</style>
