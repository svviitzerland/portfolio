<script>
	import { spring } from 'svelte/motion';

	let { children, class: className = '', ...rest } = $props();

	let buttonEl;

	const x = spring(0, { stiffness: 0.1, damping: 0.25 });
	const y = spring(0, { stiffness: 0.1, damping: 0.25 });

	function handleMouseMove(e) {
		if (!buttonEl) return;
		const rect = buttonEl.getBoundingClientRect();
		const centerX = rect.left + rect.width / 2;
		const centerY = rect.top + rect.height / 2;

		const distanceX = e.clientX - centerX;
		const distanceY = e.clientY - centerY;

		x.set(distanceX * 0.2); // Magnetic strength
		y.set(distanceY * 0.2);
	}

	function handleMouseLeave() {
		x.set(0);
		y.set(0);
	}
</script>

<div
	bind:this={buttonEl}
	onmousemove={handleMouseMove}
	onmouseleave={handleMouseLeave}
	style="transform: translate({$x}px, {$y}px)"
	class="relative inline-flex items-center justify-center transition-transform will-change-transform {className}"
	{...rest}
>
	{@render children()}
</div>
