## Description
CLI to calculate TF-IDF and download files

This CLI has the essential commands to calculate the TF-IDF weight and to download files from the internet through a given url.

To calculate the weight of a word within a given document use:

```bash 
challenge tfidf -w "ac" -f "/data/document_1.txt"
```

To download a file from the internet given the url of the file use:
```bash 
challenge wget "https://unec.edu.az/application/uploads/2014/12/pdf-sample.pdf" -d "/tmp" -f "file.pdf"
```

