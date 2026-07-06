# 📧 DNS Email Checker (Go + Web UI)
deployed URL : https://dns-email-checker.onrender.com

A simple and fast web application built in Go that checks a domain’s email DNS configuration:

- MX records (mail servers)
- SPF record (email sender policy)
- DMARC record (email authentication policy)

Includes a modern frontend UI built with HTML, CSS, and JavaScript.

---

## 🚀 Features

- 🔍 Check any domain instantly  
- 📬 Detect MX records  
- 🛡️ Validate SPF record  
- 🔐 Validate DMARC record  
- 🌐 Simple web UI  
- ⚡ Fast Go backend (`net` package)  
- 📊 Clean JSON API response  

---

## 🧠 How it works

The backend uses Go’s built-in DNS resolver:

- `net.LookupMX()` → checks mail servers (MX records)  
- `net.LookupTXT()` → checks SPF & DMARC records  

Frontend sends a request:
