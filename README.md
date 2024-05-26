
<h1 align="center" style="font-weight: bold;">📹 Vimeo Video Chunk Uploader</h1>

![go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Google Cloud](https://img.shields.io/badge/GoogleCloud-%234285F4.svg?style=for-the-badge&logo=google-cloud&logoColor=white)

<p align="center">
 <a href="#config">Configuration</a> • 
  <a href="#running">Running the Project</a> •
 <a href="#notes">Notes</a> •
 <a href="#support">Support</a>
</p>

## Project Description

This project, developed in Go 1.22, aims to automate the process of uploading videos to Vimeo. The complete workflow includes:

1. Download a video from Google Cloud Storage.
2. Split the video into 128MB chunks.
3. Create a video record via the Vimeo API.
4. Upload each of the chunks.
5. Verify the successful upload using the Vimeo API.

<h2 id="config">⚙️ Configuration</h2>

### Environment Variables

The project uses environment variables to configure necessary keys and credentials. A `.env.sample` file is provided as an example. Rename this file to `.env` and fill in the required values.

### Service Account

Depending on the permissions configured on your Google Cloud Storage bucket, you might need to place your service account JSON file in the `./config` directory.

<h2 id="running">🚀 Running the Project</h2>

### Prerequisites

* Go 1.22 installed
* Credentials configured in the `.env` file
* (Optional) Service account file in the `./config` directory

### Execution

To run the project, use the following commands:

<pre><div class="dark bg-gray-950 rounded-md border-[0.5px] border-token-border-medium"><div class="flex items-center relative text-token-text-secondary bg-token-main-surface-secondary px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span>bash</span><div class="flex items-center"><span class="" data-state="closed"></span></div></div></div></pre>

<pre><div class="dark bg-gray-950 rounded-md border-[0.5px] border-token-border-medium"><div class="overflow-y-auto p-4 text-left undefined" dir="ltr"><code class="!whitespace-pre hljs language-bash"># Navigate to the project directory
cd path/to/your/project

# Run the project
go run .
</code></div></div></pre>

### Build

To create an executable binary of the project:

<pre><div class="dark bg-gray-950 rounded-md border-[0.5px] border-token-border-medium"><div class="flex items-center relative text-token-text-secondary bg-token-main-surface-secondary px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span>bash</span><div class="flex items-center"><span class="" data-state="closed"></span></div></div></div></pre>

<pre><div class="dark bg-gray-950 rounded-md border-[0.5px] border-token-border-medium"><div class="overflow-y-auto p-4 text-left undefined" dir="ltr"><code class="!whitespace-pre hljs language-bash"># Navigate to the project directory
cd path/to/your/project

# Compile the project
go build -o vimeo-uploader
</code></div></div></pre>

### Usage

After compiling, you can run the generated binary:

<pre><div class="dark bg-gray-950 rounded-md border-[0.5px] border-token-border-medium"><div class="flex items-center relative text-token-text-secondary bg-token-main-surface-secondary px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span>bash</span><div class="flex items-center"><span class="" data-state="closed"></span></div></div></div></pre>

<pre><div class="overflow-y-auto p-4 text-left undefined" dir="ltr"><code class="!whitespace-pre hljs language-bash">./vimeo-uploader
</code></div></pre>

<h2 id="notes">📝 Notes</h2>

* Ensure all dependencies and credentials are correctly configured before running the project.
* Check the permissions of your Google Cloud Storage bucket to ensure the application has access to the videos.

<h2 id="support">📞 Support</h2>

For any questions or issues, please open an issue on [GitHub Issues](https://github.com/your-username/your-repository/issues).
