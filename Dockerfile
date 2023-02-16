FROM python:3.8-slim-buster

# Set the working directory
WORKDIR /app

# Copy the project files into the image
COPY . /app

# Install the required packages
RUN apt-get update \
    && apt-get install -y \
    build-essential \
    libasound2-dev \
    ffmpeg \
    && pip install -r requirements.txt

# Set the environment variables
ENV PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1 \
    DJANGO_SETTINGS_MODULE=dogboard.settings \
    PYTHONPATH=/app

# Expose port 8000 for Daphne
EXPOSE 8000

# Start the Django development server
CMD ["./manage.py", "runserver", "0.0.0.0:8000"]