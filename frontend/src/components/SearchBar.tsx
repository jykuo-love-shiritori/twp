import React from 'react'
import '@/components/style.css';
import '@/style/globals.css';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faSearch } from "@fortawesome/free-solid-svg-icons";

const SearchBar = () => {
    return (
        <div className="input_container">
            <input
                type="text"
                placeholder="Search"
                className="search"
            />
            <FontAwesomeIcon icon={faSearch} className="search_icon" />

        </div>
    )
}

export default SearchBar