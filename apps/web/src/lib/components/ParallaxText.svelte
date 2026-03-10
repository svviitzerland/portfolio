<script>
	import { spring } from 'svelte/motion';

	let { children, speed = 1, class: className = '' } = $props();

	let y = spring(0, { stiffness: 0.1, damping: 0.3 });
	let element;
	let scrollY = $state(0);
	let innerHeight = $state(0);

	$effect(() => {
		if (!element) return;

		const offset = scrollY * speed * 0.1;
		y.set(offset);
	});
</script>

<svelte:window bind:scrollY bind:innerHeight />

<div
	bind:this={element}
	style="transform: translateY({$y}px)"
	class="will-change-transform {className}"
>
	{@render children()}
</div>
