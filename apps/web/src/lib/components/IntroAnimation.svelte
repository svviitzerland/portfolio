<script>
	import { onMount } from 'svelte';

	let { onComplete = () => {} } = $props();

	let phase = $state(0);
	let errorLines = $state([]);
	let showCursor = $state(true);
	let corruption = $state(0);
	let deviceType = $state('windows'); // windows, mac, ios, android, linux
	let sshCopied = $state(false);

	// Boot log state for Phase 2
	let repairLines = $state([]);

	// Device-specific error messages
	const errorMessages = {
		windows: {
			face: ':(',
			title: 'Your PC ran into a problem and needs to restart.',
			subtitle: "We're just collecting some error info, and then we'll restart for you.",
			errors: [
				'*** STOP: 0x0000007E (0xFFFFFFFFC0000005)',
				'SYSTEM_THREAD_EXCEPTION_NOT_HANDLED',
				'',
				'Technical information:',
				'*** portfolio.sys - Address FFFFF880009A8D40',
				'',
				'Collecting data for crash dump...',
				'Dumping physical memory to disk: 100',
				'Physical memory dump complete.'
			],
			qrText: 'visit https://error.local/stop'
		},
		mac: {
			face: '',
			title: 'Your Mac restarted because of a problem.',
			subtitle: 'Press a key or wait a few seconds to continue starting up.',
			errors: [
				'panic(cpu 0 caller 0xffffff8000): Kernel trap at 0x0',
				'RAX: 0x0000000000000000, RBX: 0xffffff80108c6000',
				'RCX: 0x0000000000000000, RDX: 0xffffff80108c6000',
				'RSP: 0xffffff81e7ad3a30, RBP: 0xffffff81e7ad3a50',
				'',
				'Backtrace (CPU 0), Frame : Return Address',
				'0xffffff81e7ad3700 : 0xffffff8000208e26',
				'0xffffff81e7ad3750 : 0xffffff8000209872',
				'',
				'BSD process name: portfolio'
			],
			qrText: 'support.apple.com/mac/restart'
		},
		ios: {
			face: '',
			title: 'iPhone needs to restart',
			subtitle: 'An error occurred. Your device will restart automatically.',
			errors: [
				'Exception Type: EXC_BAD_ACCESS (SIGSEGV)',
				'Exception Subtype: KERN_INVALID_ADDRESS at 0x0',
				'Termination Reason: Namespace SIGNAL, Code 11',
				'',
				'Thread 0 Crashed:',
				'0   Portfolio    0x1000bfc84 main + 180',
				'1   libdyld.dylib  0x1a0f5b8f0 start + 4',
				'',
				'Binary Images:',
				'0x100000000 - 0x100ffffff Portfolio arm64'
			],
			qrText: 'support.apple.com/iphone'
		},
		android: {
			face: '',
			title: 'System UI has stopped',
			subtitle: 'Unfortunately, System UI has stopped working.',
			errors: [
				'FATAL EXCEPTION: main',
				'Process: com.portfolio.app, PID: 12847',
				'java.lang.NullPointerException',
				'   at com.portfolio.MainActivity.onCreate()',
				'   at android.app.Activity.performCreate()',
				'',
				'Build fingerprint: google/coral/coral:14',
				'',
				'Waiting for debugger...'
			],
			qrText: 'support.google.com/android'
		},
		linux: {
			face: '',
			title: 'Kernel panic - not syncing',
			subtitle: 'VFS: Unable to mount root fs on unknown-block(0,0)',
			errors: [
				'CPU: 0 PID: 1 Comm: swapper/0 Not tainted 6.1.0',
				'Hardware name: Generic PC',
				'Call Trace:',
				'  dump_stack_lvl+0x37/0x4d',
				'  panic+0x107/0x2d8',
				'  mount_block_root+0x161/0x21b',
				'',
				'Kernel Offset: disabled',
				'---[ end Kernel panic - not syncing ]---'
			],
			qrText: 'kernel.org/doc'
		}
	};

	const repairBootLines = [
		{ text: '[    0.000000] Initiating system recovery...', type: 'dim' },
		{ text: '[    0.000100] Loading portfolio kernel modules...', type: 'dim' },
		{ text: '[    0.001204] CPU: Intel Core i9-13900K @ 5.80GHz', type: 'dim' },
		{ text: '[    0.002100] Restoring filesystem...', type: 'dim' },
		{ text: '[    0.003000] Loading Svelte runtime...             [  OK  ]', type: 'ok' },
		{ text: '[    0.003500] Loading TailwindCSS...                [  OK  ]', type: 'ok' },
		{ text: '[    0.004000] Connecting PostgreSQL...               [  OK  ]', type: 'ok' },
		{ text: '[    0.004500] Starting FastAPI backend...            [  OK  ]', type: 'ok' },
		{ text: '[    0.005000] Deploying to AWS Cloudflare...         [  OK  ]', type: 'ok' },
		{ text: '[    0.005500] Mounting Docker containers...          [  OK  ]', type: 'ok' },
		{ text: '[    0.006000] portfolio.service: Active (running)', type: 'dim' },
		{ text: '[    0.006500] ✓ System restored successfully.', type: 'success' }
	];

	function detectDevice() {
		if (typeof navigator === 'undefined') return 'windows';
		const ua = navigator.userAgent.toLowerCase();
		const platform = navigator.platform?.toLowerCase() || '';

		if (/iphone|ipad|ipod/.test(ua)) return 'ios';
		if (/android/.test(ua)) return 'android';
		if (/mac/.test(platform) || /macintosh/.test(ua)) return 'mac';
		if (/linux/.test(platform) && !/android/.test(ua)) return 'linux';
		return 'windows';
	}

	// Prevent body scroll when intro is active
	$effect(() => {
		if (phase < 4) {
			document.body.style.overflow = 'hidden';
		} else {
			document.body.style.overflow = '';
		}

		return () => {
			document.body.style.overflow = '';
		};
	});

	onMount(() => {
		deviceType = detectDevice();

		const cursorInterval = setInterval(() => {
			showCursor = !showCursor;
		}, 500);

		const errors = errorMessages[deviceType].errors;
		let lineIdx = 0;
		const bsodInterval = setInterval(() => {
			if (lineIdx < errors.length) {
				errorLines = [...errorLines, errors[lineIdx]];
				lineIdx++;
				corruption = Math.min(100, corruption + 10);
			} else {
				clearInterval(bsodInterval);
				setTimeout(() => {
					phase = 1;
					setTimeout(() => {
						phase = 2;
						startTerminals();
					}, 1500);
				}, 1000);
			}
		}, 150);

		return () => {
			clearInterval(cursorInterval);
			clearInterval(bsodInterval);
		};
	});

	function startTerminals() {
		let idx = 0;
		const interval = setInterval(() => {
			if (idx < repairBootLines.length) {
				repairLines = [...repairBootLines.slice(0, idx + 1)];
				idx++;
			} else {
				clearInterval(interval);
				setTimeout(() => {
					phase = 3;
				}, 500);
			}
		}, 120);
	}

	function getBootLineClass(type) {
		if (type === 'ok') return 'boot-ok';
		if (type === 'success') return 'boot-success';
		return 'boot-dim';
	}

	let currentError = $derived(errorMessages[deviceType]);

	function openSSH() {
		// Open SSH client via protocol URI
		window.open('ssh://ssh.farhanaulianda.my.id', '_self');
		navigator.clipboard.writeText('ssh ssh.farhanaulianda.my.id');
		sshCopied = true;
		setTimeout(() => (sshCopied = false), 3000);
	}

	function skipToPortfolio() {
		phase = 4;
		onComplete();
	}
</script>

{#if phase < 4}
	<div
		class="intro-overlay {deviceType}"
		class:phase-glitch={phase === 1}
		class:phase-repair={phase === 2}
		class:phase-ssh={phase === 3}
		style="--corruption: {corruption}%"
	>
		<!-- Corruption effects layer -->
		<div class="corruption-layer">
			<div class="noise"></div>
			<div class="scanlines"></div>
			{#if phase === 1}
				<div class="screen-tear"></div>
				<div class="screen-tear delay"></div>
			{/if}
		</div>

		<!-- Error Screen -->
		{#if phase === 0 || phase === 1}
			<div class="error-screen">
				<div class="error-content">
					{#if deviceType === 'windows'}
						<div class="error-face">:(</div>
					{:else if deviceType === 'mac'}
						<div class="mac-icon">⚠</div>
					{:else if deviceType === 'ios'}
						<div class="ios-icon">
							<svg viewBox="0 0 24 24" fill="currentColor"
								><path
									d="M18.71 19.5c-.83 1.24-1.71 2.45-3.05 2.47-1.34.03-1.77-.79-3.29-.79-1.53 0-2 .77-3.27.82-1.31.05-2.3-1.32-3.14-2.53C4.25 17 2.94 12.45 4.7 9.39c.87-1.52 2.43-2.48 4.12-2.51 1.28-.02 2.5.87 3.29.87.78 0 2.26-1.07 3.81-.91.65.03 2.47.26 3.64 1.98-.09.06-2.17 1.28-2.15 3.81.03 3.02 2.65 4.03 2.68 4.04-.03.07-.42 1.44-1.38 2.83M13 3.5c.73-.83 1.94-1.46 2.94-1.5.13 1.17-.34 2.35-1.04 3.19-.69.85-1.83 1.51-2.95 1.42-.15-1.15.41-2.35 1.05-3.11z"
								/></svg
							>
						</div>
					{:else if deviceType === 'android'}
						<div class="android-icon">
							<svg viewBox="0 0 24 24" fill="currentColor"
								><path
									d="M6 18c0 .55.45 1 1 1h1v3.5c0 .83.67 1.5 1.5 1.5s1.5-.67 1.5-1.5V19h2v3.5c0 .83.67 1.5 1.5 1.5s1.5-.67 1.5-1.5V19h1c.55 0 1-.45 1-1V8H6v10zM3.5 8C2.67 8 2 8.67 2 9.5v7c0 .83.67 1.5 1.5 1.5S5 17.33 5 16.5v-7C5 8.67 4.33 8 3.5 8zm17 0c-.83 0-1.5.67-1.5 1.5v7c0 .83.67 1.5 1.5 1.5s1.5-.67 1.5-1.5v-7c0-.83-.67-1.5-1.5-1.5zm-4.97-5.84l1.3-1.3c.2-.2.2-.51 0-.71-.2-.2-.51-.2-.71 0l-1.48 1.48A5.84 5.84 0 0012 1c-.96 0-1.86.23-2.66.63L7.85.15c-.2-.2-.51-.2-.71 0-.2.2-.2.51 0 .71l1.31 1.31A5.983 5.983 0 006 7h12c0-1.99-.97-3.75-2.47-4.84zM10 5H9V4h1v1zm5 0h-1V4h1v1z"
								/></svg
							>
						</div>
					{:else}
						<div class="linux-icon">🐧</div>
					{/if}

					<div class="error-title">{currentError.title}</div>
					<div class="error-subtitle">{currentError.subtitle}</div>

					<div class="error-progress">
						<span>{Math.min(100, Math.floor(corruption))}% complete</span>
					</div>

					<div class="error-terminal">
						{#each errorLines as line, i (i)}
							<div class="error-line">{line}</div>
						{/each}
						{#if showCursor && phase === 0}
							<span class="cursor">_</span>
						{/if}
					</div>

					<div class="error-footer">
						<div class="qr-placeholder"></div>
						<div class="qr-text">
							<span>For more information:</span>
							<span>{currentError.qrText}</span>
						</div>
					</div>
				</div>
			</div>
		{/if}

		<!-- Boot Log Repair -->
		{#if phase === 2}
			<div class="boot-log">
				{#each repairLines as line, i (i)}
					<div class="boot-line {getBootLineClass(line.type)}">{line.text}</div>
				{/each}
				{#if showCursor}<span class="cursor boot-cursor">_</span>{/if}
			</div>
		{/if}

		<!-- SSH Suggestion Screen -->
		{#if phase === 3}
			<div class="ssh-screen">
				<div class="ssh-content">
					<div class="ssh-check">&#10003;</div>
					<div class="ssh-heading">System restored. Pick your interface.</div>

					<div class="ssh-choices">
						<button class="choice-card" onclick={skipToPortfolio}>
							<div class="choice-icon">
								<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
									<rect x="2" y="3" width="20" height="14" rx="2" />
									<path d="M8 21h8M12 17v4" />
								</svg>
							</div>
							<div class="choice-label">Web</div>
							<div class="choice-desc">browser</div>
						</button>

						<div class="choice-or">or</div>

						<button class="choice-card" onclick={openSSH}>
							<div class="choice-icon">
								<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
									<path d="M4 17l6-6-6-6M12 19h8" />
								</svg>
							</div>
							<div class="choice-label">SSH</div>
							<div class="choice-desc">terminal</div>
						</button>
					</div>

					<div class="ssh-command-line" class:copied={sshCopied}>
						<span class="ssh-prompt">$</span>
						<span class="ssh-text">ssh ssh.farhanaulianda.my.id</span>
						{#if sshCopied}
							<span class="ssh-badge">connecting... (copied!)</span>
						{/if}
					</div>
				</div>
			</div>
		{/if}
	</div>
{/if}

<style>
	.intro-overlay {
		position: fixed;
		inset: 0;
		z-index: 99999;
		font-family:
			'Segoe UI',
			-apple-system,
			system-ui,
			sans-serif;
		overflow: hidden;
		transition: all 0.5s ease;
	}

	/* Device-specific backgrounds */
	.intro-overlay.windows {
		background: #0078d4;
	}
	.intro-overlay.mac {
		background: #1d1d1f;
	}
	.intro-overlay.ios {
		background: #000;
	}
	.intro-overlay.android {
		background: #1a1a1a;
	}
	.intro-overlay.linux {
		background: #2e2e2e;
	}

	.intro-overlay.phase-repair {
		background: #0a0a0f !important;
	}

	.intro-overlay.phase-ssh {
		background: #09090b !important;
	}

	/* Corruption effects */
	.corruption-layer {
		position: absolute;
		inset: 0;
		pointer-events: none;
		z-index: 10;
	}

	.noise {
		position: absolute;
		inset: 0;
		opacity: calc(var(--corruption) * 0.002);
		background-image: url("data:image/svg+xml,%3Csvg viewBox='0 0 256 256' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='noise'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.8' numOctaves='4' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23noise)'/%3E%3C/svg%3E");
		animation: noise 0.2s steps(8) infinite;
	}

	@keyframes noise {
		0%,
		100% {
			transform: translate(0, 0);
		}
		25% {
			transform: translate(-5%, 5%);
		}
		50% {
			transform: translate(5%, -5%);
		}
		75% {
			transform: translate(-5%, -5%);
		}
	}

	.scanlines {
		position: absolute;
		inset: 0;
		background: repeating-linear-gradient(
			0deg,
			transparent 0px,
			transparent 1px,
			rgba(0, 0, 0, 0.1) 1px,
			rgba(0, 0, 0, 0.1) 2px
		);
	}

	.screen-tear {
		position: absolute;
		left: 0;
		right: 0;
		height: 3px;
		background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.8), transparent);
		animation: tear 0.15s linear infinite;
	}

	.screen-tear.delay {
		animation-delay: 0.075s;
	}

	@keyframes tear {
		0% {
			top: -5%;
			transform: translateX(-100%);
		}
		100% {
			top: 105%;
			transform: translateX(100%);
		}
	}

	/* Glitch phase */
	.phase-glitch {
		animation: glitch-shake 0.1s steps(2) infinite;
	}

	.phase-glitch .error-content {
		animation: glitch-text 0.15s steps(2) infinite;
	}

	@keyframes glitch-shake {
		0% {
			transform: translate(0);
			filter: hue-rotate(0deg);
		}
		25% {
			transform: translate(-3px, 2px);
		}
		50% {
			transform: translate(3px, -2px);
			filter: hue-rotate(90deg);
		}
		75% {
			transform: translate(-2px, -2px);
		}
		100% {
			transform: translate(2px, 2px);
			filter: hue-rotate(-90deg);
		}
	}

	@keyframes glitch-text {
		0% {
			text-shadow:
				-2px 0 red,
				2px 0 cyan;
		}
		50% {
			text-shadow:
				2px 0 red,
				-2px 0 cyan;
		}
		100% {
			text-shadow:
				-2px 0 red,
				2px 0 cyan;
		}
	}

	/* Error Screen */
	.error-screen {
		position: absolute;
		inset: 0;
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 24px;
	}

	@media (min-width: 600px) {
		.error-screen {
			padding: 40px;
		}
	}

	.error-content {
		max-width: 900px;
		color: white;
		overflow: hidden;
	}

	.error-face {
		font-size: clamp(80px, 15vw, 150px);
		font-weight: 100;
		margin-bottom: 24px;
	}

	.mac-icon,
	.linux-icon {
		font-size: clamp(60px, 12vw, 100px);
		margin-bottom: 24px;
	}

	.ios-icon,
	.android-icon {
		width: clamp(60px, 12vw, 100px);
		height: clamp(60px, 12vw, 100px);
		margin-bottom: 24px;
		color: white;
	}

	.ios-icon svg,
	.android-icon svg {
		width: 100%;
		height: 100%;
	}

	.error-title {
		font-size: clamp(20px, 4vw, 32px);
		font-weight: 300;
		margin-bottom: 16px;
	}

	.error-subtitle {
		font-size: clamp(14px, 2vw, 18px);
		opacity: 0.8;
		margin-bottom: 32px;
	}

	.error-progress {
		font-size: clamp(18px, 3vw, 24px);
		margin-bottom: 32px;
	}

	.error-terminal {
		font-family: 'SF Mono', 'Consolas', 'Monaco', monospace;
		font-size: clamp(10px, 1.2vw, 13px);
		line-height: 1.6;
		opacity: 0.7;
		margin-bottom: 32px;
		max-height: 200px;
		overflow: hidden;
	}

	.error-line {
		animation: fade-in 0.1s ease;
	}

	.error-footer {
		display: flex;
		align-items: center;
		gap: 16px;
		opacity: 0.6;
		font-size: 12px;
	}

	.qr-placeholder {
		width: 80px;
		height: 80px;
		background: white;
		border-radius: 4px;
		flex-shrink: 0;
	}

	.qr-text {
		display: flex;
		flex-direction: column;
		gap: 4px;
	}

	/* Boot Log */
	.boot-log {
		position: absolute;
		inset: 0;
		padding: 40px;
		font-family: 'SF Mono', 'Monaco', 'Consolas', 'Courier New', monospace;
		font-size: clamp(12px, 1.4vw, 15px);
		line-height: 1.8;
		overflow: hidden;
		animation: fade-in 0.5s ease;
		display: flex;
		flex-direction: column;
		justify-content: center;
	}

	@media (max-width: 600px) {
		.boot-log {
			padding: 24px 16px;
			font-size: 10px;
			line-height: 1.7;
		}
	}

	.boot-line {
		animation: slide-in 0.1s ease;
		white-space: pre-wrap;
		word-break: break-all;
	}

	.boot-dim {
		color: #71717a;
	}

	.boot-ok {
		color: #9ca3af;
	}

	.boot-success {
		color: #10b981;
		font-weight: 600;
	}

	.boot-cursor {
		color: #71717a;
		margin-left: 0;
	}

	.cursor {
		display: inline-block;
		animation: blink 0.8s step-end infinite;
	}

	@keyframes blink {
		50% {
			opacity: 0;
		}
	}

	@keyframes fade-in {
		from {
			opacity: 0;
		}
		to {
			opacity: 1;
		}
	}

	@keyframes slide-in {
		from {
			opacity: 0;
			transform: translateX(-10px);
		}
		to {
			opacity: 1;
			transform: translateX(0);
		}
	}

	/* SSH Choice Screen */
	.ssh-screen {
		position: absolute;
		inset: 0;
		display: flex;
		align-items: center;
		justify-content: center;
		animation: fade-in 0.6s ease;
	}

	.ssh-content {
		text-align: center;
		color: #e4e4e7;
	}

	.ssh-check {
		font-size: 48px;
		color: #10b981;
		margin-bottom: 24px;
		animation: scale-in 0.4s cubic-bezier(0.34, 1.56, 0.64, 1);
	}

	@keyframes scale-in {
		from {
			transform: scale(0);
			opacity: 0;
		}
		to {
			transform: scale(1);
			opacity: 1;
		}
	}

	.ssh-heading {
		font-family: 'SF Mono', 'Monaco', 'Consolas', monospace;
		font-size: clamp(16px, 2.5vw, 22px);
		color: #a1a1aa;
		margin-bottom: 40px;
	}

	.ssh-choices {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 24px;
		margin-bottom: 36px;
	}

	.choice-card {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 12px;
		background: #18181b;
		border: 1px solid #27272a;
		border-radius: 16px;
		padding: 32px 48px;
		cursor: pointer;
		transition: all 0.3s ease;
		min-width: 160px;
	}

	.choice-card:hover {
		border-color: #10b981;
		box-shadow: 0 0 30px rgba(16, 185, 129, 0.12);
		transform: translateY(-4px);
	}

	.choice-card:active {
		transform: translateY(-2px);
	}

	.choice-icon {
		width: 36px;
		height: 36px;
		color: #71717a;
		transition: color 0.3s ease;
	}

	.choice-card:hover .choice-icon {
		color: #10b981;
	}

	.choice-label {
		font-family: 'SF Mono', 'Monaco', 'Consolas', monospace;
		font-size: 20px;
		font-weight: 600;
		color: #e4e4e7;
	}

	.choice-desc {
		font-family: 'SF Mono', 'Monaco', 'Consolas', monospace;
		font-size: 12px;
		color: #52525b;
	}

	.choice-or {
		font-family: 'SF Mono', 'Monaco', 'Consolas', monospace;
		font-size: 13px;
		color: #3f3f46;
	}

	.ssh-command-line {
		display: inline-flex;
		align-items: center;
		gap: 8px;
		font-family: 'SF Mono', 'Monaco', 'Consolas', monospace;
		font-size: 13px;
		color: #3f3f46;
		transition: color 0.2s ease;
	}

	.ssh-command-line.copied {
		color: #10b981;
	}

	.ssh-prompt {
		color: #52525b;
	}

	.ssh-text {
		color: inherit;
	}

	.ssh-badge {
		font-size: 11px;
		color: #10b981;
		animation: fade-in 0.2s ease;
	}

	@media (max-width: 600px) {
		.choice-card {
			padding: 24px 32px;
			min-width: 120px;
		}
		.ssh-choices {
			gap: 16px;
		}
	}
</style>
