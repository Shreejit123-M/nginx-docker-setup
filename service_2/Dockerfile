# Use a lightweight Python image
FROM python:3.9-slim

# Set working directory inside container
WORKDIR /app

# Copy all files from current directory
COPY . .

# Install Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Expose the app port
EXPOSE 8002

# Run the app
CMD ["python", "app.py"]
