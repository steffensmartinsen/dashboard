import { Input, InputGroup,  } from '@chakra-ui/react'
import React from 'react'

function FootballInput(props) {

    return (
        <InputGroup size='md' className={props.className}>
            <Input
                pr='4.5rem'
                value={props.team}
                type={"text"}
                placeholder='Choose team'
                onChange={props.onChange}
            />
        </InputGroup>
    )
}

export default FootballInput