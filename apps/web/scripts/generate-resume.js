import fs from 'fs';
import path from 'path';
import handlebars from 'handlebars';
import puppeteer from 'puppeteer';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

async function generateResume() {
    try {
        console.log('Starting resume generation...');

        // 1. Read the JSON data
        const cvPath = path.join(__dirname, '../src/lib/data/cv.json');
        if (!fs.existsSync(cvPath)) {
            throw new Error('cv.json not found at ' + cvPath);
        }
        const cvData = JSON.parse(fs.readFileSync(cvPath, 'utf8'));
        console.log('Loaded cv.json');

        // 2. Read the Template
        const templatePath = path.join(__dirname, 'templates', 'resume.hbs');
        const templateSource = fs.readFileSync(templatePath, 'utf8');

        // 3. Read the CSS
        const cssPath = path.join(__dirname, 'templates', 'style.css');
        const css = fs.readFileSync(cssPath, 'utf8');

        // 4. Compile Handlebars
        const template = handlebars.compile(templateSource);
        let html = template({ ...cvData }); // Don't pass CSS here

        // Manual CSS Injection
        html = html.replace('<!-- CSS_INJECT -->', `<style>${css}</style>`);

        // 5. Generate PDF with Puppeteer
        console.log('Launching browser...');
        const browser = await puppeteer.launch({
            args: ['--no-sandbox', '--disable-setuid-sandbox'],
            headless: 'new'
        });
        const page = await browser.newPage();

        // Set content
        await page.setContent(html, { waitUntil: 'networkidle0' });

        // Ensure output directory exists
        const outputDir = path.join(__dirname, '../static');
        if (!fs.existsSync(outputDir)) {
            fs.mkdirSync(outputDir, { recursive: true });
        }

        // Create PDF
        const outputPath = path.join(outputDir, 'resume.pdf');
        await page.pdf({
            path: outputPath,
            format: 'A4',
            printBackground: true,
            margin: {
                top: '8mm',
                bottom: '8mm',
                left: '0mm',
                right: '0mm'
            }
        });

        console.log(`Resume generated successfully at: ${outputPath}`);

        await browser.close();

    } catch (error) {
        console.error('Error generating resume:', error);
        process.exit(1);
    }
}

generateResume();
