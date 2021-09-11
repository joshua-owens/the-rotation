<?php

namespace App;

interface Scraper {
    public function scrape(): Recipe;
}
