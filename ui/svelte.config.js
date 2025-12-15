import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://svelte.dev/docs/kit/integrations
	// for more information about preprocessors
	preprocess: vitePreprocess(),

	kit: {
		// adapter-auto only supports some environments, see https://svelte.dev/docs/kit/adapter-auto for a list.
		// If your environment is not supported, or you settled on a specific environment, switch out the adapter.
		// See https://svelte.dev/docs/kit/adapters for more information about adapters.
		adapter: adapter({
			// SvelteKit 默认输出到 build 目录 (纯 Vite 是 dist，这里要注意区别)
			pages: 'build',
			assets: 'build',
			fallback: 'index.html', // ⚠️ 关键：生成 index.html 供 Go 后端做 SPA 路由回退
			precompress: false,
			strict: true
		})
	}
};

export default config;
