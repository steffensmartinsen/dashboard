import { Input, InputGroup,  } from '@chakra-ui/react'
import React from 'react'

function EmailInput(props) {

    return (
        <InputGroup size='md' className={props.className}>
            <Input
                pr='4.5rem'
                value={props.email}
                type={"email"}
                placeholder='Enter e-mail'
                onChange={props.onChange}
            />
        </InputGroup>
    )
}

export default EmailInput