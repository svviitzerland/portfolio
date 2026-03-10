<script>
	import { fly } from 'svelte/transition';

	let { label, onclick, children, x = 0, y = 0, onmove } = $props();

	let dragging = $state(false);
	let dragOffsetX = $state(0);
	let dragOffsetY = $state(0);
	let startPointer = { x: 0, y: 0 };
	let hasMoved = false;

	const DRAG_THRESHOLD = 5;
	const GRID_SIZE = 20;

	// During drag, show dragged position; otherwise show prop position
	let displayX = $derived(dragging ? x + dragOffsetX : x);
	let displayY = $derived(dragging ? y + dragOffsetY : y);

	function snapToGrid(value) {
		return Math.round(value / GRID_SIZE) * GRID_SIZE;
	}

	function onPointerDown(e) {
		if (e.button !== 0) return;

		dragging = true;
		hasMoved = false;
		dragOffsetX = 0;
		dragOffsetY = 0;
		startPointer = { x: e.clientX, y: e.clientY };

		e.currentTarget.setPointerCapture(e.pointerId);
		e.preventDefault();
	}

	function onPointerMove(e) {
		if (!dragging) return;

		const dx = e.clientX - startPointer.x;
		const dy = e.clientY - startPointer.y;

		if (!hasMoved && Math.abs(dx) < DRAG_THRESHOLD && Math.abs(dy) < DRAG_THRESHOLD) {
			return;
		}

		hasMoved = true;
		dragOffsetX = dx;
		dragOffsetY = dy;
	}

	function onPointerUp(e) {
		if (!dragging) return;
		dragging = false;

		if (hasMoved) {
			const finalX = snapToGrid(x + dragOffsetX);
			const finalY = snapToGrid(y + dragOffsetY);
			onmove?.(finalX, finalY);
		} else {
			e.stopPropagation();
			onclick?.();
		}

		dragOffsetX = 0;
		dragOffsetY = 0;
	}
</script>

<div
	class="absolute"
	style="left: {displayX}px; top: {displayY}px; z-index: {dragging ? 50 : 10};"
	in:fly={{ y: 10, duration: 300 }}
>
	<button
		class="group flex flex-col items-center gap-1.5 w-[76px] p-2 rounded-lg hover:bg-white/10 transition-all duration-200 cursor-default select-none"
		style="touch-action: none;"
		onpointerdown={onPointerDown}
		onpointermove={onPointerMove}
		onpointerup={onPointerUp}
		onpointercancel={onPointerUp}
	>
		<!-- Icon container -->
		<div
			class="w-14 h-14 flex items-center justify-center rounded-[14px] bg-zinc-800/80 border border-zinc-700/60 shadow-lg group-hover:scale-105 group-hover:border-zinc-600 transition-all duration-200"
		>
			{@render children()}
		</div>

		<span
			class="text-[11px] text-white/90 text-center leading-tight font-medium max-w-full truncate"
			style="text-shadow: 0 1px 4px rgba(0,0,0,0.9), 0 0 10px rgba(0,0,0,0.5);"
		>
			{label}
		</span>
	</button>
</div>
