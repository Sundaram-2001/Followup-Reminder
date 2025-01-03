import adapter from '@sveltejs/adapter-auto';

/** @type {import('@sveltejs/kit').Config} */
const config = {
  kit: {
    adapter: adapter(),
  },
  vite: {
    server: {
      proxy: {
        '/email': 'http://localhost:8090'
      }
    }
  }
};

export default config;