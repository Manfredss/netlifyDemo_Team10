import logging

from flask import Flask, request, redirect, render_template, url_for
import sqlite3

app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html')