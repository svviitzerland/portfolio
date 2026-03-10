<script>
	import cv from '$lib/data/cv.json';
</script>

<div class="resume-container">
	<!-- Header -->
	<header class="resume-header">
		<h1 class="name">{cv.basics.name}</h1>
		<p class="label">{cv.basics.label}</p>
	</header>

	<!-- Contact Box -->
	<div class="contact-box">
		<div class="contact-item">
			<svg class="contact-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
				/>
			</svg>
			<span>{cv.basics.email}</span>
		</div>
		<div class="contact-item">
			<svg class="contact-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"
				/>
			</svg>
			<span>{cv.basics.phone}</span>
		</div>
		<div class="contact-item">
			<svg class="contact-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9a9 9 0 01-9-9m9 9c1.657 0 3-4.03 3-9s-1.343-9-3-9m0 18c-1.657 0-3-4.03-3-9s1.343-9 3-9m-9 9a9 9 0 019-9"
				/>
			</svg>
			<span>{cv.basics.website}</span>
		</div>
	</div>

	<!-- Summary -->
	<section class="section summary-section">
		<p>{cv.basics.summary}</p>
	</section>

	<!-- Skills -->
	<section class="section">
		<div class="section-heading">
			<h2>Skills</h2>
			<div class="line"></div>
		</div>
		<div class="skills-grid">
			{#each cv.skills as skill (skill)}
				<div class="skill-item">
					<span class="bullet">•</span>
					{skill}
				</div>
			{/each}
		</div>
	</section>

	<!-- Achievements -->
	<section class="section">
		<div class="section-heading">
			<h2>Achievements</h2>
			<div class="line"></div>
		</div>
		<div class="timeline">
			{#each cv.achievements as achievement (achievement.title + achievement.date)}
				<div class="timeline-item">
					<div class="timeline-marker"></div>
					<div class="timeline-content">
						<div class="timeline-header">
							<div class="title-company"><strong>{achievement.title}</strong></div>
							<div class="date">{achievement.date}</div>
						</div>
						{#if achievement.summary}<div class="job-summary">{achievement.summary}</div>{/if}
						{#if achievement.highlights && achievement.highlights.length > 0}
							<ul class="highlights-list">
								{#each achievement.highlights as highlight (highlight)}
									<li>{highlight}</li>
								{/each}
							</ul>
						{/if}
					</div>
				</div>
			{/each}
		</div>
	</section>

	<!-- Job Experience -->
	<section class="section">
		<div class="section-heading">
			<h2>Job Experience</h2>
			<div class="line"></div>
		</div>
		<div class="timeline">
			{#each cv.work as job (job.company + job.position + job.startDate)}
				<div class="timeline-item">
					<div class="timeline-marker"></div>
					<div class="timeline-content">
						<div class="timeline-header">
							<div class="title-company">
								<strong>{job.position}, {job.company}</strong>
							</div>
							<div class="date">{job.startDate} - {job.endDate}</div>
						</div>
						{#if job.summary}<div class="job-summary">{job.summary}</div>{/if}
						{#if job.highlights && job.highlights.length > 0}
							<div class="responsibilities-label">Responsibilities:</div>
							<ul class="highlights-list">
								{#each job.highlights as highlight (highlight)}
									<li>{highlight}</li>
								{/each}
							</ul>
						{/if}
						{#if job.stack}
							<div class="stack-list">
								<strong>Stack:</strong>
								<div class="stack-text">{job.stack}</div>
							</div>
						{/if}
					</div>
				</div>
			{/each}
		</div>
	</section>

	<!-- Projects -->
	<section class="section">
		<div class="section-heading">
			<h2>Projects</h2>
			<div class="line"></div>
		</div>
		<div class="timeline">
			{#each cv.projects as project (project.name + project.date)}
				<div class="timeline-item">
					<div class="timeline-marker"></div>
					<div class="timeline-content">
						<div class="timeline-header">
							<div class="title-company">
								<strong>{project.name}</strong>{#if project.url}
									<span class="project-url">({project.url})</span>{/if}
							</div>
							<div class="date">
								{project.date}
								{#if project.endDate}- {project.endDate}{/if}
							</div>
						</div>
						{#if project.summary}<div class="job-summary">{project.summary}</div>{/if}
						{#if project.highlights && project.highlights.length > 0}
							<div class="responsibilities-label">Responsibilities:</div>
							<ul class="highlights-list">
								{#each project.highlights as highlight (highlight)}
									<li>{highlight}</li>
								{/each}
							</ul>
						{/if}
						{#if project.stack}
							<div class="stack-list">
								<strong>Stack:</strong>
								<div class="stack-text">{project.stack}</div>
							</div>
						{/if}
					</div>
				</div>
			{/each}
		</div>
	</section>

	<!-- Education -->
	<section class="section">
		<div class="section-heading">
			<h2>Education</h2>
			<div class="line"></div>
		</div>
		<div class="timeline">
			{#each cv.education as edu (edu.institution + edu.startDate)}
				<div class="timeline-item">
					<div class="timeline-marker"></div>
					<div class="timeline-content">
						<div class="timeline-header">
							<div class="title-school">
								<span class="school"><strong>{edu.institution}</strong></span>
							</div>
							<div class="date">{edu.startDate} - {edu.endDate}</div>
						</div>
						<div class="degree">( {edu.studyType} ) {edu.area}</div>
					</div>
				</div>
			{/each}
		</div>
	</section>

	<!-- Organization -->
	{#if cv.experiences_in_organization && cv.experiences_in_organization.length > 0}
		<section class="section">
			<div class="section-heading">
				<h2>Experiences In Organization</h2>
				<div class="line"></div>
			</div>
			<div class="timeline">
				{#each cv.experiences_in_organization as org (org.organization + org.position + org.startDate)}
					<div class="timeline-item">
						<div class="timeline-marker"></div>
						<div class="timeline-content">
							<div class="timeline-header">
								<div class="title-company">
									<strong>{org.position} of {org.organization}</strong>
								</div>
								<div class="date">{org.startDate} - {org.endDate}</div>
							</div>
							{#if org.summary}<p class="job-summary">{org.summary}</p>{/if}
						</div>
					</div>
				{/each}
			</div>
		</section>
	{/if}

	<!-- Talks -->
	{#if cv.talks && cv.talks.length > 0}
		<section class="section">
			<div class="section-heading">
				<h2>Talks & Presentations</h2>
				<div class="line"></div>
			</div>
			<div class="timeline">
				{#each cv.talks as talk (talk.title + talk.date)}
					<div class="timeline-item">
						<div class="timeline-marker"></div>
						<div class="timeline-content">
							<div class="timeline-header">
								<div class="title-company"><strong>{talk.title}</strong></div>
								<div class="date">{talk.date}</div>
							</div>
							<div class="job-summary">{talk.event}</div>
							{#if talk.summary}<p class="job-summary">{talk.summary}</p>{/if}
						</div>
					</div>
				{/each}
			</div>
		</section>
	{/if}

	<!-- Certifications -->
	<section class="section">
		<div class="section-heading">
			<h2>Certifications</h2>
			<div class="line"></div>
		</div>
		<div class="timeline">
			{#each cv.certifications as cert (cert.name + cert.date)}
				<div class="timeline-item">
					<div class="timeline-marker"></div>
					<div class="timeline-content">
						<div class="timeline-header">
							<div class="title-company"><strong>{cert.name}</strong></div>
							<div class="date">{cert.date}</div>
						</div>
						{#if cert.score}<div class="job-summary">Score: {cert.score}</div>{/if}
					</div>
				</div>
			{/each}
		</div>
	</section>

	<!-- Socials -->
	<section class="section">
		<div class="section-heading">
			<h2>Social Links</h2>
			<div class="line"></div>
		</div>
		<div class="social-list">
			{#each cv.socials as social (social.network)}
				<div class="social-item">
					<span class="social-label">{social.network}:</span>
					<a
						href={social.url}
						target="_blank"
						class="social-url"
						style="word-break: break-all; color: inherit; text-decoration: none;">{social.url}</a
					>
				</div>
			{/each}
		</div>
	</section>
</div>

<style>
	:root {
		--primary-color: #6d5e41;
		--header-color: #5e4b35;
		--accent-color: #8e806a;
		--separator-color: #a09581;
		--text-color: #333;
		--text-dark: #444;
		--text-medium: #555;
		--text-secondary: #666;
		--bg-color: #fff;
		--bg-secondary: #f9f8f6;
		--timeline-line-color: #dcd6c9;
		--border-color: #6d5e41;
	}

	.resume-container {
		font-family: 'Merriweather', serif;
		color: var(--text-color);
		line-height: 1.4;
		background-color: #fff;
		font-size: 10pt;
		width: 100%;
		max-width: 210mm;
		min-height: 297mm;
		margin: 0 auto;
		padding: 30px 40px;
		position: relative;
		padding-bottom: 50px;
		box-sizing: border-box;
		overflow-x: hidden;
	}

	.resume-container * {
		box-sizing: border-box;
		margin: 0;
		padding: 0;
	}

	.resume-header {
		text-align: left;
		margin-bottom: 5px;
	}

	.name {
		font-family: 'Merriweather', serif;
		font-size: 32px;
		font-weight: 700;
		color: var(--header-color);
		margin-bottom: 5px;
	}

	.label {
		font-family: 'Merriweather', serif;
		font-style: italic;
		font-size: 14px;
		color: var(--text-dark);
		margin-bottom: 15px;
	}

	.contact-box {
		display: flex;
		justify-content: space-between;
		align-items: center;
		border: 1px solid var(--accent-color);
		padding: 8px 20px;
		background-color: var(--bg-secondary);
		margin-bottom: 25px;
		flex-wrap: wrap;
		gap: 15px;
	}

	.contact-item {
		font-family: 'Merriweather', serif;
		font-size: 10pt;
		color: var(--primary-color);
		display: flex;
		align-items: center;
		gap: 8px;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
		min-width: 0;
		flex: 0 1 auto;
	}

	.contact-item span {
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
		min-width: 0;
	}

	.contact-icon {
		width: 14px;
		height: 14px;
		flex-shrink: 0;
		color: var(--primary-color);
	}

	.section {
		margin-bottom: 15px;
	}

	.section-heading {
		display: flex;
		align-items: center;
		margin-bottom: 15px;
	}

	.section-heading h2 {
		font-family: 'Merriweather', serif;
		font-size: 16px;
		font-weight: 700;
		color: var(--header-color);
		margin-right: 15px;
		white-space: nowrap;
	}

	.section-heading .line {
		height: 2px;
		background-color: var(--separator-color);
		width: 100%;
		margin-top: 4px;
	}

	.summary-section p {
		text-align: justify;
		line-height: 1.5;
		font-size: 10pt;
		color: #333;
	}

	.timeline {
		position: relative;
		padding-left: 5px;
	}

	.timeline-item {
		position: relative;
		padding-left: 20px;
		margin-bottom: 15px;
		break-inside: avoid;
	}

	.timeline::before {
		content: '';
		position: absolute;
		left: 9.5px;
		top: 5px;
		bottom: 5px;
		width: 1px;
		background-color: var(--timeline-line-color);
	}

	.timeline-marker {
		position: absolute;
		left: 0;
		top: 5px;
		width: 10px;
		height: 10px;
		border-radius: 50%;
		background-color: var(--accent-color);
		z-index: 1;
	}

	.timeline-header {
		display: flex;
		justify-content: space-between;
		align-items: flex-start;
		margin-bottom: 2px;
		gap: 10px;
		flex-wrap: wrap;
	}

	.title-school,
	.title-company {
		font-size: 11pt;
		color: #333;
		flex: 1;
		min-width: 0;
		overflow-wrap: break-word;
		word-break: break-word;
	}

	.title-company strong,
	.title-school strong {
		color: var(--text-dark);
		font-weight: 700;
	}

	.date {
		font-size: 10pt;
		font-style: italic;
		color: var(--text-secondary);
		white-space: nowrap;
	}

	.degree {
		font-size: 10pt;
		color: var(--text-medium);
	}

	.job-summary {
		font-size: 10pt;
		margin-top: 2px;
		margin-bottom: 4px;
		color: #333;
	}

	.responsibilities-label {
		font-size: 10pt;
		margin-bottom: 2px;
		color: var(--text-dark);
	}

	.highlights-list {
		list-style-type: disc;
		margin-left: 20px;
		font-size: 10pt;
		color: #333;
	}

	.highlights-list li {
		margin-bottom: 2px;
	}

	.skills-grid {
		display: grid;
		grid-template-columns: repeat(4, 1fr);
		gap: 5px;
		padding-left: 10px;
		width: 100%;
		overflow: hidden;
	}

	.skill-item {
		font-size: 10pt;
		color: var(--text-dark);
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	.bullet {
		color: var(--text-dark);
		font-weight: bold;
		margin-right: 5px;
	}

	.stack-list {
		margin-top: 5px;
		font-size: 10pt;
		font-style: italic;
	}

	.stack-text {
		display: inline;
	}

	.social-list {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 8px;
		padding-left: 10px;
	}

	.social-item {
		font-size: 10pt;
		display: flex;
		gap: 8px;
	}

	.social-label {
		font-weight: 600;
		color: var(--text-dark);
		min-width: 70px;
	}

	.social-url {
		color: var(--text-medium);
	}

	.project-url {
		font-size: 9pt;
		font-weight: 400;
		color: var(--text-secondary);
		font-style: italic;
	}

	@media print {
		.resume-container {
			width: 100%;
			margin: 0;
			box-shadow: none;
			padding: 0 30px;
		}
	}

	/* Custom Scrollbar */
	.resume-container::-webkit-scrollbar {
		width: 6px;
	}

	.resume-container::-webkit-scrollbar-track {
		background: transparent;
	}

	.resume-container::-webkit-scrollbar-thumb {
		background-color: var(--header-color);
		border-radius: 3px;
	}

	.resume-container::-webkit-scrollbar-thumb:hover {
		background-color: var(--primary-color);
	}
</style>
