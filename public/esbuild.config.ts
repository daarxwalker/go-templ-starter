import esbuild from 'esbuild'
import tailwindPlugin from 'esbuild-plugin-tailwindcss'
import manifestPlugin from 'esbuild-plugin-manifest'
import { cleanPlugin } from 'esbuild-clean-plugin'

try {
	await esbuild.build({
		entryPoints: ['./public/index.ts'],
		assetNames: '[name]-[hash]',
		entryNames: 'bundle-[hash]',
		outdir: './public/static/dist',
		bundle: true,
		plugins: [
			tailwindPlugin({}),
			manifestPlugin({
				shortNames: true,
			}),
			cleanPlugin({}),
		],
	})
	console.log('assets built successfully!')
} catch(e) {
	console.error(e)
}