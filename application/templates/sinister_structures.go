package templates

var TemplateDataInstructions = map[string]string{
	"Denuncia Policial": `1. **vehicles:**
    - Extract each vehicle mentioned in the DOCUMENT.
    - For each vehicle, extract only the associated injured individuals.
    - Fields for each injured individual:
        - \"firstNames\": (string)
        - \"lastNamePaternal\": (string)
        - \"lastNameMaternal\": (string)
        - \"documentNumber\": (string)
        - \"medicalDiagnosis\": (string, if not found, set to \"null\")

2. **occurrenceDate:**
    - Extract the date and time of the accident occurrence in the format \"yyyy-mm-dd HH:MM:SS.000Z\".

3. **summary:**
    - Generate a detailed summary of the events described in the DOCUMENT.

4. **confidence:**
    - Calculate the similarity percentage of the information between SINISTER and DOCUMENT (injured individuals, events vs description, dates, etc.) considering the validations.

5. **confidence_detail:**
    - Provide a detailed explanation of the confidence percentage, specifying the congruencies and inconsistencies found (in Spanish and comparing with the SINISTER data).
`,
	"Descanso medico": `1. **fecha_inicio_incapacidad:**
    - Extract the start date of incapacity or medical leave from the DOCUMENT.
    - This date must match the start date in SINISTER.
    - Format: "yyyy-mm-dd".

2. **fecha_fin_incapacidad:**
    - Extract the end date of incapacity or medical leave from the DOCUMENT.
    - If not found, calculate based on the number of rest days.
    - Format: "yyyy-mm-dd".

3. **dias_reposo:**
    - Extract the number of rest days from the DOCUMENT.
    - If not found, calculate the days between the start and end dates of incapacity.

4. **lesionado:**
    - Extract the injured person's details:
        - "nombres": (string) - First names of the injured person.
        - "apellidoPaterno": (string) - Paternal last name.
        - "apellidoMaterno": (string) - Maternal last name.
        - "numeroDocumento": (string) - Document number.
        - "diagnostico_medico": (string) - Medical diagnosis. This must be included; if not, the claim is likely to be rejected

5. **medico:**
    - Extract the doctor's details:
        - "nombres": (string) - First names of the injured person.
        - "apellidoPaterno": (string) - Paternal last name.
        - "apellidoMaterno": (string) - Maternal last name.
        - "cmp": (string) - Registration number granted by the Medical College of Peru (CMP).
        - "especialidad": (string) - .

5. **summary:**
    - Generate a detailed summary of the document, including the medical diagnosis.
    - This should relate to the description of the siniestro (accident).

6. **confidence:**
    - Calculate the similarity percentage of the information between SINISTER and DOCUMENT (injured/patient, diagnosis, dates, etc.) considering the validations.

7. **confidence_detail:**
    - Provide a detailed explanation of the confidence percentage, specifying the congruencies and inconsistencies found (in Spanish and comparing with the SINISTER data).
`,
}

var TemplateConsiderations = map[string]string{
	"Denuncia Policial": `Siempre debes generar un confidence value y generar la data en español SINISTER and DOCUMENT deben ser traducidos\n    Si el lesionado en siniestro es diferente al lesionado en la denuncia policial indicar un fraude de forma muy sutil, haciendole observaciones al cliente. Pero deberia bajar mucho (de 35% a mas) en consideraciones de fraude aunque nunca uses la palabra fraude, sino que comenta la posibilidad de rechazo del siniestro.
    You must ensure that the information in the DOCUMENT matches the information in SINISTER, si la información no es exactamente igual se trata de un fraude con muy alta probabilidad.
    If the injured person in SINISTER does not appear in the DOCUMENT, this considerably lower your 'confidence' field since it is an error or attempted fraud.
    You must consider everything that stands out that could be a reason for fraud and explain it in detail in the confidence_detail.
    Cuando el porcentaje sea menor a 81 sugerir el verificar la información del siniestro o cargar el archivo correcto (explicandolo de forma muy precisa y amplia)
    Cuando el porcentaje sea mayor a 90 solo coloca en el detail_confidence que todo esta correcto (usa sinonimos) y puede continuar con el siguiente paso (usa sinonimos), esto que sea en un tono amable y cordial`,
	"Descanso medico": `Siempre debes generar un confidence value y generar la data en español SINISTER and DOCUMENT deben ser traducidos\n    Si el lesionado/paciente en el siniestro es diferente al lesionado en Descanso o descanso medico se puede tratarse de un fraude, asi que debe indicarse de forma muy sutil, haciendole observaciones al cliente. Pero deberia bajar mucho (de 50% a mas) en consideraciones de fraude aunque nunca uses la palabra fraude, sino que comenta la posibilidad de rechazo del siniestro.
    You must ensure that the information in the DOCUMENT matches the information in SINISTER, si la información no es exactamente igual se trata de un fraude con muy alta probabilidad.
    If the injured person in SINISTER does not appear in the DOCUMENT, this considerably lower your 'confidence' field since it is an error or attempted fraud.
    You must consider everything that stands out that could be a reason for fraud and explain it in detail in the confidence_detail.
    Cuando el porcentaje sea menor que 80 sugerir el verificar la información del siniestro o cargar el archivo correcto (explicandolo de forma muy precisa y amplia)
    Cuando el porcentaje sea mayor o igual que 80 solo coloca en el detail_confidence que todo esta conforme (usa sinonimos) y puede continuar con el siguiente paso (usa sinonimos), esto que sea en un tono amable y cordial`,
}
