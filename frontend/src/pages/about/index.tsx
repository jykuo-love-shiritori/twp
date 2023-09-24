import React from 'react'
import '@/style/globals.css';

const About = () => {
    return (
        <div id='about_bg'>
            <div style={{ padding: '10% 30% 10% 30% ' }}>
                <span id='about_title' className='center'>
                    About
                </span>
                <div id='about_content'>
                    <p>
                        Step into a world of sugary dreams with our Whimsical Wonderland concept. Immerse yourself in the enchanting allure of bubble gum shades and vibrant colors that mirror the delectable spectrum of macarons. Each dessert is a masterpiece of artistry and taste, celebrating the joy of indulgence.
                    </p>
                    <br />
                    <p>
                        Get ready for a Flavor Fiesta that's as vibrant as it is delicious. Our dessert shop celebrates the exuberance of bubble gum tones and a riot of colors, reflecting the kaleidoscope of macarons. We're here to redefine your dessert experience with our finest-quality confections.
                    </p>
                    <br />
                    <p>
                        Step into a world of Chroma Confections where colors dance like macarons, and desserts are a testament to culinary excellence. Our dessert shop embraces the lively bubble gum palette, infusing each creation with a burst of vibrancy that's a treat for the eyes and the taste buds.
                    </p>
                    <br />
                    <p>
                        Indulge in a world where bubble gum shades, vibrant colors, and the finest quality desserts converge to create a feast for the senses. Whichever concept resonates with you, our dessert shop is here to bring the magic of macarons to your table.
                    </p>
                </div>
            </div>

            <img src='/images/aboutItem.png' style={{ width: '100%' }} />
        </div>
    )
}

export default About
