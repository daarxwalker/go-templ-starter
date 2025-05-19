import esbuild from 'esbuild'
import tailwindPlugin from 'esbuild-plugin-tailwindcss'
import manifestPlugin from 'esbuild-plugin-manifest'
import { cleanPlugin } from 'esbuild-clean-plugin'
import type { Plugin } from 'esbuild';

const buildWatcherLoggerPlugin: Plugin = {
	name: 'build-watcher-logger',
	setup(build) {
		let startTime: number
		build.onStart(() => {
			startTime = performance.now()
			const now = new Date()
			const time = now.toLocaleTimeString()
			console.log(`building assets...`)
		})
		build.onEnd((result) => {
			const endTime = performance.now()
			const duration = endTime - startTime
		  if (result.errors.length > 0) {
		    console.error(`assets build failed in ${duration.toFixed(2)}ms!`)
			  for (const error of result.errors) {
			    console.error(error.text)
			  }
		  } else {
		    console.log(`assets built successfully in ${duration.toFixed(2)}ms!`)
		  }
		});
	},
};

try {
	let ctx = await esbuild.context({
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
			buildWatcherLoggerPlugin,
		],
	})
	await ctx.watch()
	console.log('watching assets...')
} catch(e) {
	console.error(e)
}