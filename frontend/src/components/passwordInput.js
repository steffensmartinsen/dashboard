import { Input, InputGroup, InputRightElement, Button } from '@chakra-ui/react'
import React from 'react'

function PasswordInput(props) {
    const [show, setShow] = React.useState(false)
    const handleClick = () => setShow(!show)

    return (
        <InputGroup size='md' className={props.className}>
            <Input
                pr='4.5rem'
                value={props.password}
                type={show ? 'text' : 'password'}
                placeholder={props.placeholder}
                onChange={props.onChange}
            />
            <InputRightElement width='4.5rem'>
                <Button h='1.75rem' size='sm' onClick={handleClick}>
                    {show ? 'Hide' : 'Show'}
                </Button>
            </InputRightElement>
        </InputGroup>
    )
}

export default PasswordInput