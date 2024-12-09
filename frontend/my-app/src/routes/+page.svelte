<script>
	import { z } from 'zod';

	let email = '';
	let numberOfDays = '';
	let error = '';

	const schema = z.object({
		email: z.string().email('Invalid email address'),
		numberOfDays: z.string().regex(/^\d+$/, 'Days must be a number')
	});

	function handleSubmit() {
		try {
			const validatedData = schema.parse({ email, numberOfDays });
			// Submit logic here
			console.log('Submitted:', validatedData);
			error = '';
		} catch (err) {
			error = err.errors[0].message;
		}
	}
</script>

<div class="container">
	<form on:submit|preventDefault={handleSubmit}>
		<div class="form-group">
			<label for="email">Email</label>
			<input 
				type="text" 
				id="email" 
				bind:value={email} 
				placeholder="Enter your email"
			/>
		</div>
		<div class="form-group">
			<label for="days">Number of Days</label>
			<input 
				type="text" 
				id="days" 
				bind:value={numberOfDays} 
				placeholder="Enter number of days"
			/>
		</div>
		{#if error}
			<div class="error">{error}</div>
		{/if}
		<button type="submit">Submit</button>
	</form>
</div>

<style>
	.container {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100vh;
	}

	form {
		width: 300px;
		padding: 20px;
		border: 1px solid #ccc;
		border-radius: 8px;
		box-shadow: 0 2px 4px rgba(0,0,0,0.1);
	}

	.form-group {
		margin-bottom: 15px;
	}

	label {
		display: block;
		margin-bottom: 5px;
	}

	input {
		width: 100%;
		padding: 8px;
		border: 1px solid #ddd;
		border-radius: 4px;
	}

	.error {
		color: red;
		margin-bottom: 10px;
	}

	button {
		width: 100%;
		padding: 10px;
		background-color: #007bff;
		color: white;
		border: none;
		border-radius: 4px;
		cursor: pointer;
	}

	button:hover {
		background-color: #0056b3;
	}
</style>