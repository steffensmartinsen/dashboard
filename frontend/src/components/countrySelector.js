import React, { useState, useMemo } from 'react'
import Select from 'react-select'
import countryList from 'react-select-country-list'

function CountrySelector(props) {
    const { country, setCountry, className } = props
    const options = useMemo(() => countryList().getData(), [])

    const changeHandler = value => {
        setCountry(value)
    }

    return <Select options={options} value={country} onChange={changeHandler} className={className} placeholder={country}/>
}

export default CountrySelector