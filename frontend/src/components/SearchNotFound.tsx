import React from 'react'
import SearchBar from './SearchBar'

const SearchNotFound = () => {
    return (
        <div className='center pureBG flex-wrapper' style={{ padding: '20% 10% 10% 10%' }}>
            <div className='search_not_found'>Search not found</div>
            <div className='white_word'>please try again</div>
            <br />
            <SearchBar />
        </div>
    )
}

export default SearchNotFound;
