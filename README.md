# Study Topics Randomizer (Go CLI)

## Overview

This is a small **command-line tool written in Go** that randomly selects a study topic each time it is run. The purpose of the project is to support **consistent daily learning** by removing decision fatigue and encouraging focused exploration of technical subjects.

Rather than choosing *what* to study each day, this tool chooses for me.

---

## Motivation

Maintaining consistent study habits can be difficult when there are too many possible topics to choose from. This project was created to:

- Encourage daily learning
- Reduce time spent deciding what to study
- Promote exploration across a broad range of topics
- Serve as a lightweight, reusable learning aid

On some days, the selected topic may lead to:
- Reading documentation or articles  
- Reviewing fundamentals  
- Designing or building a small project to explore the topic in depth  

---

## How It Works

- Study topics are stored in Go data structures (arrays / slices)
- Each run of the CLI randomly selects one topic
- The output is intentionally simple and fast

The project is designed to be **easy to extend**, allowing new topics or entire subject groups to be added with minimal effort.

---

## Extensibility

This project is intentionally structured to evolve over time. Planned and ongoing enhancements include:

- Expanding existing topic lists
- Adding multiple subject categories (e.g., Networking, Security, Go, Cloud)
- Introducing flags or arguments to filter topics by category
- Improving output formatting or metadata
- Adding persistence or tracking in the future if useful

The tool grows alongside my learning goals.

---

## Why Go?

Go was chosen for this project because it offers:

- Fast startup time for CLI tools
- Simple syntax and strong standard library support
- Easy distribution as a single binary
- Practical relevance for systems and tooling work

---

## Usage

```bash
go run main.go
```
Or, once built:

```bash
./study-randomizer
```

Each execution returns a single randomized study topic.

Purpose of This Repository
This project is both:

A practical personal utility

A demonstration of how I approach small, purpose-driven tooling

It reflects an emphasis on:

Intentional design

Simplicity

Continuous improvement

Future Ideas
Weighted topic selection

Daily topic history

Exporting results to a file

Integration with task or note-taking tools
