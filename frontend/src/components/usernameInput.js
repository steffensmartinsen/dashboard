import { Input, InputGroup,  } from '@chakra-ui/react'
import React from 'react'

function UsernameInput(props) {

    return (
        <InputGroup size='md'>
            <Input
                pr='4.5rem'
                value={props.username}
                type={"text"}
                placeholder='Enter username'
                onChange={props.onChange}
            />
        </InputGroup>
    )
}

export default UsernameInput