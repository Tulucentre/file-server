# Tulucentre File Server

A simple HTTP file server written in Go for uploading, downloading, and deleting files via REST API endpoints.

## Features

- Upload files to custom folder paths (auto-creates directories)
- Download files by path
- Delete files by path
- Simple API endpoints

## API Endpoints

### Upload File

- **POST** `/api/addfile`
- Form fields:
  - `file`: File to upload (multipart/form-data)
  - `path`: Destination path (e.g., `folder/subfolder/filename.txt`)
- Creates any missing directories in the path.

### Download File

- **GET** `/api/getfile?path=your/file/path.txt`
- Returns the file as a download.

### Delete File

- **DELETE** `/api/delete?path=your/file/path.txt`
- Deletes the specified file.

## Getting Started

### Prerequisites

- Go 1.18 or newer
- `.env` file

### Development

```sh
make dev
```

### Build

```sh
make build
```

### Run

```sh
make start
```

The server will start on `http://localhost:8080`.
