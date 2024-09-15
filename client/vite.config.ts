import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [react()],
	server: {
		host: '0.0.0.0',
		port: 3000,
		watch: {
			usePolling: true,
		},
	},
	// We must determine in which port the app will run when
	// we execute pnpm run preview
	// host true will expose the project in public addresses
	preview: {
		host: true,
		port: 8080,
	},
});
