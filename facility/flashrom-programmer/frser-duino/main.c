/*
 * This file is part of the frser-duino project.
 *
 * Copyright (C) 2010 Urja Rannikko <urjaman@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301 USA
 */

#include "main.h"
#include "uart.h"
#include "frser.h"


int main(void) {
#if (defined __AVR_ATmega1280__)||(defined __AVR_ATmega1281__)||(defined __AVR_ATmega2560__)||(defined __AVR_ATmega2561__)
	/* Arduino Mega: LED off (but driven) for now */
	PORTB &= ~_BV(7);
	DDRB |= _BV(7);
#endif
	uart_init();
	frser_main();
}


